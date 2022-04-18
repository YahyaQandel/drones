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
const BATTERY_LEVEL_SCHEDULER_WAIT_SCEONDS = 20
const SAVE_LOG_SCHEDULER_WAIT_SCEONDS = 60

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
	droneLogRepository := repository.NewDroneLogRepository(dbClient)
	droneUsecase := usecase.NewDroneUsecase(droneRepository, medicationRepository)
	droneActionUsecase := usecase.NewDroneActionUsecase(droneRepository, medicationRepository, droneMedicationRepository)
	droneLogUsecase := usecase.NewLogUsecase(droneRepository, droneLogRepository)
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
	schedulerUsecase := usecase.NewSchedulerUsecase(droneActionUsecase, droneLogUsecase, droneRepository)
	go runSchedulerEveryXSeconds(schedulerUsecase.UpdateLoadedDronesBatteryLevel, BATTERY_LEVEL_SCHEDULER_WAIT_SCEONDS)
	go runSchedulerEveryXSeconds(droneLogUsecase.LogDroneInfo, SAVE_LOG_SCHEDULER_WAIT_SCEONDS)
	server.StartServer(apis)
}

func runSchedulerEveryXSeconds(schedluerFn fn, seconds int) {
	s := gocron.NewScheduler()
	s.Every(uint64(seconds)).Seconds().Do(schedluerFn, context.Background())
	<-s.Start()
}
