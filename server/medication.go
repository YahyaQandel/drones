package server

import (
	"fmt"
	"log"
	"net/http"

	"drones.com/usecase"
)

type MedicationApi struct {
	medicationUsecase usecase.IMedicationeUsecase
}

func NewMedicationAPI(useCase usecase.IMedicationeUsecase) MedicationApi {
	return MedicationApi{
		medicationUsecase: useCase,
	}
}

func (api MedicationApi) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response, err := api.medicationUsecase.RegisterMedication(ctx, r)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
