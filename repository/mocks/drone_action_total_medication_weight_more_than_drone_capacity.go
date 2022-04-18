package mocks

import (
	"context"

	"drones.com/repository"
	"drones.com/repository/entity"
)

type MockedDroneActionTotalMedicationsWeightMoreThanDroneCapacityRepository struct {
}

func NewMockedDroneActionTotalMedicationsWeightMoreThanDroneCapacityRepository() repository.IDroneActionRepo {
	return &MockedDroneActionTotalMedicationsWeightMoreThanDroneCapacityRepository{}
}

func (cdb MockedDroneActionTotalMedicationsWeightMoreThanDroneCapacityRepository) CreateDroneMedication(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error) {
	return entity.DroneMedication{}, nil
}

func (cdb MockedDroneActionTotalMedicationsWeightMoreThanDroneCapacityRepository) Get(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error) {
	return entity.DroneMedication{}, nil
}

func (cdb MockedDroneActionTotalMedicationsWeightMoreThanDroneCapacityRepository) IsNotFoundErr(err error) bool {
	return true
}

func (cdb MockedDroneActionTotalMedicationsWeightMoreThanDroneCapacityRepository) GetDroneMedications(ctx context.Context, droneMedication entity.DroneMedication) ([]entity.DroneMedication, error) {
	return []entity.DroneMedication{
		{
			DroneSerialNumber: "xxx",
			MedicationCode:    "PD1",
		},
		{
			DroneSerialNumber: "xxx",
			MedicationCode:    "PD2",
		},
		{
			DroneSerialNumber: "xxx",
			MedicationCode:    "PD3",
		},
	}, nil
}
