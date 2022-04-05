package server

import (
	"net/http"

	"drones.com/usecase"
)

type DroneApi struct {
	droneUsecase usecase.DroneUsecase
}

func NewDroneAPI(useCase usecase.DroneUsecase) DroneApi {
	return DroneApi{
		droneUsecase: useCase,
	}
}

func (api DroneApi) Create(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte{})
}
