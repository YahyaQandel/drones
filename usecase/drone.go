package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"drones.com/repository"
	repoEntity "drones.com/repository/entity"
	"drones.com/usecase/entity"
	"github.com/asaskevich/govalidator"
)

type DroneUsecase interface {
	RegisterDrone(ctx context.Context, request []byte) ([]byte, error)
}

type droneUsecase struct {
	droneRepo repository.IDroneRepo
}

func NewDroneUsecase(droneRepository repository.IDroneRepo) DroneUsecase {
	return droneUsecase{droneRepo: droneRepository}
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
	// by default drone battery capacity is 100 and state is IDLE in creation
	droneRequest.BatteryCapacity = 100
	droneRequest.State = string(entity.IDLE)
	droneRepoEntity := repoEntity.Drone{}
	droneRepoEntity.SerialNumber = droneRequest.SerialNumber
	droneRepoEntity.Model = droneRequest.Model
	droneRepoEntity.Weight = droneRequest.Weight
	droneRepoEntity.BatteryCapacity = droneRequest.Weight
	droneRepoEntity.State = droneRequest.State
	repoQueryResult, err := d.droneRepo.Create(ctx, droneRepoEntity)
	if err != nil {
		return []byte{}, err
	}
	result, err := json.Marshal(repoQueryResult)
	if err != nil {
		return []byte{}, nil
	}
	return result, nil
}

func IsInvalidWeightErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Drone.weight of type float64")
}

func IsValidDroneModel(requestDrone entity.Drone) bool {
	for i := 0; i < len(entity.DroneModels); i++ {
		if requestDrone.Model == string(entity.DroneModels[i]) {
			return true
		}
	}
	return false
}
