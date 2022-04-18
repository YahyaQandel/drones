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
	GetLoadedMedicationItems(ctx context.Context, request []byte) ([]repoEntity.Medication, error)
	GetAvailableDrones(ctx context.Context) (drones []repoEntity.Drone, err error)
	GetLoadedDrones(ctx context.Context) (drones []repoEntity.Drone, err error)
	GetDroneBatteryLevel(ctx context.Context, request []byte) (batteryLevel entity.GetDroneBatteryLevelResponse, err error)
}

type droneActionUsecase struct {
	droneRepo       repository.IDroneRepo
	medicationRepo  repository.IMedicationRepo
	droneActionRepo repository.IDroneActionRepo
}

func NewDroneActionUsecase(droneRepository repository.IDroneRepo, medicationRepository repository.IMedicationRepo, droneActionRepo repository.IDroneActionRepo) IDroneActionUsecase {
	return droneActionUsecase{droneRepo: droneRepository, medicationRepo: medicationRepository, droneActionRepo: droneActionRepo}
}

func (d droneActionUsecase) LoadDrone(ctx context.Context, request []byte) (response []byte, err error) {

	loadDrone := entity.LoadDroneRequest{}
	err = json.Unmarshal(request, &loadDrone)
	if err != nil {
		return []byte{}, err
	}
	validateLoadDroneRequest, err := govalidator.ValidateStruct(loadDrone)
	if err != nil && !validateLoadDroneRequest {
		return []byte{}, err
	}
	loadDroneRepoEntity := repoEntity.DroneMedication{DroneSerialNumber: loadDrone.DroneSerialNumber, MedicationCode: loadDrone.MedicationCode}
	droneMedicationExists, err := d.droneActionRepo.Get(ctx, loadDroneRepoEntity)
	if !d.droneActionRepo.IsNotFoundErr(err) {
		return []byte{}, errors.New(fmt.Sprintf("drone with serial number '%s' is already loaded with medication with code '%s'", droneMedicationExists.DroneSerialNumber, droneMedicationExists.MedicationCode))
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
	// handle drone should not be loaded with weight greater than its capacity
	droneLoadedMedications, err := d.droneActionRepo.GetDroneMedications(ctx, loadDroneRepoEntity)
	if err != nil {
		return []byte{}, err
	}
	totalDroneWeightLoad := 0.0
	for _, droneMedication := range droneLoadedMedications {
		loadedMedication, err := d.medicationRepo.Get(ctx, repoEntity.Medication{Code: droneMedication.MedicationCode})
		if err != nil {
			return []byte{}, err
		}
		totalDroneWeightLoad += loadedMedication.Weight
	}
	if totalDroneWeightLoad+medication.Weight > drone.Weight {
		return []byte{}, errors.New(fmt.Sprintf("cannot load drone with medication as medication weight is '%0.2f' and drone is already loaded with '%0.2f'", medication.Weight, totalDroneWeightLoad))
	}

	if drone.BatteryCapacity < 25.00 {
		return []byte{}, errors.New(fmt.Sprintf("cannot load drone with medication while battery capacity is '%0.2f'", drone.BatteryCapacity))
	}
	droneMedication := repoEntity.DroneMedication{DroneSerialNumber: drone.SerialNumber, MedicationCode: medication.Code}
	_, err = d.droneActionRepo.CreateDroneMedication(ctx, droneMedication)
	if err != nil {
		return []byte{}, err
	}
	drone.State = string(entity.LOADED)
	d.droneRepo.Update(ctx, drone)
	responseMessage := []byte(fmt.Sprintf(`{"message": "drone '%s' loaded with medication '%s'"}`, drone.SerialNumber, medication.Code))
	return responseMessage, nil
}

func (d droneActionUsecase) GetLoadedMedicationItems(ctx context.Context, request []byte) (medications []repoEntity.Medication, err error) {
	getLoadedMedicationItems := entity.GetLoadedMedicationItemsRequest{}
	err = json.Unmarshal(request, &getLoadedMedicationItems)
	if err != nil {
		return []repoEntity.Medication{}, err
	}
	if err != nil {
		return []repoEntity.Medication{}, err
	}
	validateGetLoadedMedicationItemsRequest, err := govalidator.ValidateStruct(getLoadedMedicationItems)
	if err != nil && !validateGetLoadedMedicationItemsRequest {
		return []repoEntity.Medication{}, err
	}
	droneMedications, err := d.droneActionRepo.GetDroneMedications(ctx, repoEntity.DroneMedication{DroneSerialNumber: getLoadedMedicationItems.DroneSerialNumber})
	if err != nil {
		return []repoEntity.Medication{}, err
	}
	medications = []repoEntity.Medication{}
	for _, droneMedication := range droneMedications {
		medication, err := d.medicationRepo.Get(ctx, repoEntity.Medication{Code: droneMedication.MedicationCode})
		if err != nil {
			return []repoEntity.Medication{}, err
		}
		medications = append(medications, medication)
	}
	return
}

func (d droneActionUsecase) GetAvailableDrones(ctx context.Context) (drones []repoEntity.Drone, err error) {
	drones, err = d.droneRepo.GetAvailable(ctx)
	if err != nil {
		return []repoEntity.Drone{}, err
	}
	return
}

func (d droneActionUsecase) GetLoadedDrones(ctx context.Context) (drones []repoEntity.Drone, err error) {
	drones, err = d.droneRepo.GetLoaded(ctx)
	if err != nil {
		return []repoEntity.Drone{}, err
	}
	return
}

func (d droneActionUsecase) GetDroneBatteryLevel(ctx context.Context, request []byte) (batteryLevel entity.GetDroneBatteryLevelResponse, err error) {
	getDroneBatteryLevel := entity.GetDroneBatteryLevelRequest{}
	err = json.Unmarshal(request, &getDroneBatteryLevel)
	if err != nil {
		return entity.GetDroneBatteryLevelResponse{}, err
	}
	if err != nil {
		return entity.GetDroneBatteryLevelResponse{}, err
	}
	validateGetDroneBatteryLevelRequest, err := govalidator.ValidateStruct(getDroneBatteryLevel)
	if err != nil && !validateGetDroneBatteryLevelRequest {
		return entity.GetDroneBatteryLevelResponse{}, err
	}
	drone, err := d.droneRepo.Get(ctx, repoEntity.Drone{SerialNumber: getDroneBatteryLevel.DroneSerialNumber})
	if err != nil && d.droneRepo.IsNotFoundErr(err) {
		return entity.GetDroneBatteryLevelResponse{}, errors.New(fmt.Sprintf("drone not found with serial number '%s'", getDroneBatteryLevel.DroneSerialNumber))
	}
	if err != nil {
		return entity.GetDroneBatteryLevelResponse{}, err
	}

	batteryLevel.BatteryLevel = drone.BatteryCapacity
	return
}
