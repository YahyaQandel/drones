package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type MockedDroneActionAlreadyLoadedWithSameDroneAndMedicationRepository struct {
}

func NewMockedDroneActionAlreadyLoadedWithSameDroneAndMedicationRepository() repository.IDroneActionRepo {
	return &MockedDroneActionAlreadyLoadedWithSameDroneAndMedicationRepository{}
}

func (cdb MockedDroneActionAlreadyLoadedWithSameDroneAndMedicationRepository) CreateDroneMedication(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error) {
	return entity.DroneMedication{}, nil
}

func (cdb MockedDroneActionAlreadyLoadedWithSameDroneAndMedicationRepository) Get(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error) {
	return entity.DroneMedication{DroneSerialNumber: "XDX", MedicationCode: "RX"}, nil
}

func (cdb MockedDroneActionAlreadyLoadedWithSameDroneAndMedicationRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (cdb MockedDroneActionAlreadyLoadedWithSameDroneAndMedicationRepository) GetDroneMedications(ctx context.Context, droneMedication entity.DroneMedication) ([]entity.DroneMedication, error) {
	return []entity.DroneMedication{}, nil
}
