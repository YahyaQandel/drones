package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"drones.com/usecase"
)

type DroneApi struct {
	droneUsecase usecase.IDroneUsecase
}

func NewDroneAPI(useCase usecase.IDroneUsecase) DroneApi {
	return DroneApi{
		droneUsecase: useCase,
	}
}

func (api DroneApi) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}
	response, err := api.droneUsecase.RegisterDrone(ctx, requestByte)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
