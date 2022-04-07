package main

import (
	"fmt"
	"log"

	"drones.com/repository"
	"drones.com/repository/db"
	"drones.com/server"
	"drones.com/usecase"
)

func main() {
	fmt.Println("Hello World, Drones waiting... !")
	dbClient, err := db.ConnectToDatabase()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	droneRepository := repository.NewDroneRepository(dbClient)
	droneUsecase := usecase.NewDroneUsecase(droneRepository)
	medicationUsecase := usecase.NewMedicationUsecase(droneRepository)
	droneApi := server.NewDroneAPI(droneUsecase)
	medicationApi := server.NewMedicationAPI(medicationUsecase)
	apis := server.APIs{
		DronesApi:     droneApi,
		MedicationApi: medicationApi,
	}
	server.StartServer(apis)
}
