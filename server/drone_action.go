package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"drones.com/usecase"
)

type DroneActionApi struct {
	droneActionUsecase usecase.IDroneActionUsecase
}

func NewDroneActionAPI(useCase usecase.IDroneActionUsecase) DroneActionApi {
	return DroneActionApi{
		droneActionUsecase: useCase,
	}
}

func (api DroneActionApi) Load(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}
	response, err := api.droneActionUsecase.LoadDrone(ctx, requestByte)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(response)
}
