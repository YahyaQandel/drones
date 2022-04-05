package main

import (
	"fmt"

	"drones.com/server"
	"drones.com/usecase"
)

func main() {
	fmt.Println("Hello World, Drones waiting... !")
	droneUsecase := usecase.NewDroneUsecase()
	droneApi := server.NewDroneAPI(droneUsecase)
	apis := server.APIs{
		DronesApi: droneApi,
	}
	server.StartServer(apis)
}
