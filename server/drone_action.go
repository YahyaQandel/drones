package server

import (
	"encoding/json"
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

func (api DroneActionApi) GetLoadedMedicationItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params, ok := r.URL.Query()["drone_serial_number"]

	if !ok || len(params[0]) < 1 {
		errorMessage := "drone serial number is missing"
		log.Println("ERROR: ", errorMessage)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, errorMessage)))
		return
	}
	requestByte := []byte(fmt.Sprintf(`{"drone_serial_number":"%s"}`, params[0]))
	response, err := api.droneActionUsecase.GetLoadedMedicationItems(ctx, requestByte)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(response)
}

func (api DroneActionApi) GetAvailableDrones(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	drones, err := api.droneActionUsecase.GetAvailableDrones(ctx)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}
	response, err := json.Marshal(drones)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(response)
}
