package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type MockedDroneActionGetAllMedicationRepository struct {
}

func NewMockedDroneActionGetAllMedicationRepository() repository.IDroneActionRepo {
	return &MockedDroneActionGetAllMedicationRepository{}
}

func (cdb MockedDroneActionGetAllMedicationRepository) CreateDroneMedication(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error) {
	return entity.DroneMedication{}, nil
}

func (cdb MockedDroneActionGetAllMedicationRepository) Get(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error) {
	return entity.DroneMedication{}, nil
}

func (cdb MockedDroneActionGetAllMedicationRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (cdb MockedDroneActionGetAllMedicationRepository) GetDroneMedications(ctx context.Context, droneMedication entity.DroneMedication) ([]entity.DroneMedication, error) {
	return []entity.DroneMedication{
		{
			DroneSerialNumber: "XDX",
			MedicationCode:    "RX",
		},
		{
			DroneSerialNumber: "XDX",
			MedicationCode:    "UX",
		},
		{
			DroneSerialNumber: "XDX",
			MedicationCode:    "DX",
		},
	}, nil
}
