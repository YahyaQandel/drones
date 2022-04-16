package main

import (
	"context"
	"fmt"
	"log"

	"drones.com/iofile"
	"drones.com/repository"
	"drones.com/repository/db"
	"drones.com/server"
	"drones.com/usecase"
	"github.com/jasonlvhit/gocron"
)

const MEDICATION_IMAGES_PATH = "medications-images/"
const SCHEDULER_WAIT_SCEONDS = 5

type fn func(ctx context.Context) error

func main() {
	fmt.Println("Hello World, Drones waiting... !")
	dbClient, err := db.ConnectToDatabase()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	droneRepository := repository.NewDroneRepository(dbClient)
	medicationRepository := repository.NewMedicationRepository(dbClient)
	droneMedicationRepository := repository.NewDroneActionRepository(dbClient)
	droneUsecase := usecase.NewDroneUsecase(droneRepository, medicationRepository)
	droneActionUsecase := usecase.NewDroneActionUsecase(droneRepository, medicationRepository, droneMedicationRepository)
	ioFile := iofile.NewIOFile(MEDICATION_IMAGES_PATH)
	medicationUsecase := usecase.NewMedicationUsecase(medicationRepository, ioFile)
	droneApi := server.NewDroneAPI(droneUsecase)
	medicationApi := server.NewMedicationAPI(medicationUsecase)
	droneActionApi := server.NewDroneActionAPI(droneActionUsecase)
	apis := server.APIs{
		DronesApi:      droneApi,
		MedicationApi:  medicationApi,
		DroneActionApi: droneActionApi,
	}
	schedulerUsecase := usecase.NewSchedulerUsecase(droneActionUsecase, droneRepository)
	go runSchedulerEveryXSeconds(schedulerUsecase.UpdateLoadedDronesBatteryLevel)
	server.StartServer(apis)
}

func runSchedulerEveryXSeconds(schedluerFn fn) {
	log.Println("asdfasdf")
	s := gocron.NewScheduler()
	s.Every(SCHEDULER_WAIT_SCEONDS).Seconds().Do(schedluerFn, context.Background())
	<-s.Start()
}
