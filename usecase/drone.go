package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"drones.com/usecase/entity"
	"github.com/asaskevich/govalidator"
)

type DroneUsecase interface {
	RegisterDrone(ctx context.Context, request []byte) ([]byte, error)
}

type droneUsecase struct {
}

func NewDroneUsecase() DroneUsecase {
	return droneUsecase{}
}

func (d droneUsecase) RegisterDrone(ctx context.Context, request []byte) (response []byte, err error) {
	droneRequest := entity.Drone{}
	err = json.Unmarshal(request, &droneRequest)
	if IsInvalidWeightErr(err) {
		return []byte{}, errors.New(`weight: invalid input , should be float`)
	}
	if err != nil {
		return []byte{}, err
	}
	if !IsValidDroneModel(droneRequest) {
		return []byte{}, errors.New(fmt.Sprintf(`model: should be option of %v`, entity.DroneModels))
	}
	validateDroneRequest, err := govalidator.ValidateStruct(droneRequest)
	if err != nil && !validateDroneRequest {
		return []byte{}, err
	}
	response, err = json.Marshal(droneRequest)
	if err != nil && !validateDroneRequest {
		return []byte{}, err
	}
	return
}

func IsInvalidWeightErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Drone.weight of type float64")
}

func IsValidDroneModel(requestDrone entity.Drone) bool {
	for i := 0; i < len(entity.DroneModels); i++ {
		if requestDrone.Model == entity.DroneModels[i] {
			return true
		}
	}
	return false
}
