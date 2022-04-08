package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type MockedDroneActionMedicationWeightGreaterThanDroneWeightRepository struct {
}

func NewMockedDroneActionMedicationWeightGreaterThanDroneWeightRepository(client *gorm.DB) repository.IDroneActionRepo {
	return &MockedDroneActionMedicationWeightGreaterThanDroneWeightRepository{}
}

func (cdb MockedDroneActionMedicationWeightGreaterThanDroneWeightRepository) CreateDroneMedication(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error) {
	return entity.DroneMedication{}, nil
}

func (cdb MockedDroneActionMedicationWeightGreaterThanDroneWeightRepository) Get(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error) {
	return entity.DroneMedication{}, nil
}

func (cdb MockedDroneActionMedicationWeightGreaterThanDroneWeightRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
