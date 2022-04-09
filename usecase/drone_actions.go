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
	GetLoadedMedicationItems(ctx context.Context, request []byte) ([]byte, error)
	GetAvailableDrones(ctx context.Context) (drones []repoEntity.Drone, err error)
}

type droneActionUsecase struct {
	droneRepo           repository.IDroneRepo
	medicationRepo      repository.IMedicationRepo
	droneMedicationRepo repository.IDroneActionRepo
}

func NewDroneActionUsecase(droneRepository repository.IDroneRepo, medicationRepository repository.IMedicationRepo, droneMedication repository.IDroneActionRepo) IDroneActionUsecase {
	return droneActionUsecase{droneRepo: droneRepository, medicationRepo: medicationRepository, droneMedicationRepo: droneMedication}
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
	droneMedicationExists, err := d.droneMedicationRepo.Get(ctx, loadDroneRepoEntity)
	if err != nil {
		return []byte{}, err
	}
	if droneMedicationExists.DroneSerialNumber != "" {
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
	if drone.BatteryCapacity < 25.00 {
		return []byte{}, errors.New(fmt.Sprintf("cannot load drone with medication while battery capacity is '%0.2f'", drone.BatteryCapacity))
	}
	droneMedication := repoEntity.DroneMedication{DroneSerialNumber: drone.SerialNumber, MedicationCode: medication.Code}
	_, err = d.droneMedicationRepo.CreateDroneMedication(ctx, droneMedication)
	if err != nil {
		return []byte{}, err
	}
	drone.State = string(entity.LOADED)
	d.droneRepo.Update(ctx, drone)
	responseMessage := []byte(fmt.Sprintf(`{"message": "drone '%s' loaded with medication '%s'"}`, drone.SerialNumber, medication.Code))
	return responseMessage, nil
}

func (d droneActionUsecase) GetLoadedMedicationItems(ctx context.Context, request []byte) (response []byte, err error) {
	getLoadedMedicationItems := entity.GetLoadedMedicationItemsRequest{}
	err = json.Unmarshal(request, &getLoadedMedicationItems)
	if err != nil {
		return []byte{}, err
	}
	if err != nil {
		return []byte{}, err
	}
	validateGetLoadedMedicationItemsRequest, err := govalidator.ValidateStruct(getLoadedMedicationItems)
	if err != nil && !validateGetLoadedMedicationItemsRequest {
		return []byte{}, err
	}
	droneMedications, err := d.droneMedicationRepo.GetDroneMedications(ctx, repoEntity.DroneMedication{DroneSerialNumber: getLoadedMedicationItems.DroneSerialNumber})
	if err != nil {
		return []byte{}, err
	}
	medications := []repoEntity.Medication{}
	for _, droneMedication := range droneMedications {
		medication, err := d.medicationRepo.Get(ctx, repoEntity.Medication{Code: droneMedication.MedicationCode})
		if err != nil {
			return []byte{}, err
		}
		medications = append(medications, medication)
	}
	responseByte, err := json.Marshal(medications)
	if err != nil {
		return []byte{}, err
	}
	return responseByte, nil
}

func (d droneActionUsecase) GetAvailableDrones(ctx context.Context) (drones []repoEntity.Drone, err error) {
	drones, err = d.droneRepo.GetAvailable(ctx)
	if err != nil {
		return []repoEntity.Drone{}, err
	}
	if err != nil {
		return []repoEntity.Drone{}, err
	}
	return
}
