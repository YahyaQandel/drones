package main

import (
	"fmt"
	"log"

	"drones.com/iofile"
	"drones.com/repository"
	"drones.com/repository/db"
	"drones.com/server"
	"drones.com/usecase"
)

const MEDICATION_IMAGES_PATH = "medications-images/"

func main() {
	fmt.Println("Hello World, Drones waiting... !")
	dbClient, err := db.ConnectToDatabase()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	droneRepository := repository.NewDroneRepository(dbClient)
	droneUsecase := usecase.NewDroneUsecase(droneRepository)
	ioFile := iofile.NewIOFile(MEDICATION_IMAGES_PATH)
	medicationUsecase := usecase.NewMedicationUsecase(droneRepository, ioFile)
	droneApi := server.NewDroneAPI(droneUsecase)
	medicationApi := server.NewMedicationAPI(medicationUsecase)
	apis := server.APIs{
		DronesApi:     droneApi,
		MedicationApi: medicationApi,
	}
	server.StartServer(apis)
}
