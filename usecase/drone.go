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

type IDroneUsecase interface {
	RegisterDrone(ctx context.Context, request []byte) (repoEntity.Drone, error)
}

type droneUsecase struct {
	droneRepo      repository.IDroneRepo
	medicationRepo repository.IMedicationRepo
}

func NewDroneUsecase(droneRepository repository.IDroneRepo, medicationRepository repository.IMedicationRepo) IDroneUsecase {
	return droneUsecase{droneRepo: droneRepository, medicationRepo: medicationRepository}
}

func (d droneUsecase) RegisterDrone(ctx context.Context, request []byte) (response repoEntity.Drone, err error) {
	droneRequest := entity.Drone{}
	err = json.Unmarshal(request, &droneRequest)
	if IsInvalidWeightErr(err) {
		return repoEntity.Drone{}, errors.New(`weight: invalid input , should be float`)
	}
	if err != nil {
		return repoEntity.Drone{}, err
	}
	if !IsValidDroneModel(droneRequest) {
		return repoEntity.Drone{}, errors.New(fmt.Sprintf(`model: should be option of %v`, entity.DroneModels))
	}
	validateDroneRequest, err := govalidator.ValidateStruct(droneRequest)
	if err != nil && !validateDroneRequest {
		return repoEntity.Drone{}, err
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
	response, err = d.droneRepo.Create(ctx, droneRepoEntity)
	if err != nil {
		return repoEntity.Drone{}, err
	}
	return
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
