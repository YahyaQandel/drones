package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"drones.com/repository"
	repoEntity "drones.com/repository/entity"
	"drones.com/usecase/entity"
	"github.com/asaskevich/govalidator"
)

type IDroneActionUsecase interface {
	LoadDrone(ctx context.Context, request []byte) ([]byte, error)
}

type droneActionUsecase struct {
	droneRepo       repository.IDroneRepo
	medicationRepo  repository.IMedicationRepo
	droneMedication repository.IDroneActionRepo
}

func NewDroneActionUsecase(droneRepository repository.IDroneRepo, medicationRepository repository.IMedicationRepo, droneMedication repository.IDroneActionRepo) IDroneActionUsecase {
	return droneActionUsecase{droneRepo: droneRepository, medicationRepo: medicationRepository, droneMedication: droneMedication}
}

func (d droneActionUsecase) LoadDrone(ctx context.Context, request []byte) (response []byte, err error) {
	loadDrone := entity.LoadDrone{}
	err = json.Unmarshal(request, &loadDrone)
	if err != nil {
		return []byte{}, err
	}
	if err != nil {
		return []byte{}, err
	}
	validateLoadDroneRequest, err := govalidator.ValidateStruct(loadDrone)
	if err != nil && !validateLoadDroneRequest {
		return []byte{}, err
	}
	drone, err := d.droneRepo.Get(ctx, repoEntity.Drone{SerialNumber: loadDrone.DroneSerialNumber})
	if err != nil && d.droneRepo.IsNotFoundErr(err) {
		return []byte{}, errors.New(fmt.Sprintf("drone not found with serial number '%s'", loadDrone.DroneSerialNumber))
	}
	medication, err := d.medicationRepo.Get(ctx, repoEntity.Medication{Code: loadDrone.MedicationCode})
	if err != nil && d.medicationRepo.IsNotFoundErr(err) {
		return []byte{}, errors.New(fmt.Sprintf("medication not found with code '%s'", loadDrone.MedicationCode))
	}
	if medication.Weight > drone.Weight {
		return []byte{}, errors.New(fmt.Sprintf("drone weight '%0.2f' is not enough to carry medication with weight '%0.2f'", drone.Weight, medication.Weight))
	}
	if drone.BatteryCapacity < 25.00 {
		return []byte{}, errors.New(fmt.Sprintf("cannot load drone with medication while battery capacity is '%0.2f'", drone.BatteryCapacity))
	}
	droneMedication := repoEntity.DroneMedication{DroneSerialNumber: drone.SerialNumber, MedicationCode: medication.Code}
	_, err = d.droneMedication.CreateDroneMedication(ctx, droneMedication)
	if err != nil {
		return []byte{}, err
	}
	responseMessage := []byte(fmt.Sprintf(`{"message": "drone '%s' loaded with medication '%s'"}`, drone.SerialNumber, medication.Code))
	return responseMessage, nil
}
