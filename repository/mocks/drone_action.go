package mocks

import (
	"context"

	"drones.com/repository"
	"drones.com/repository/entity"
)

type MockedDroneActionRepository struct {
}

func NewMockedDroneActionRepository() repository.IDroneActionRepo {
	return &MockedDroneActionRepository{}
}

func (cdb MockedDroneActionRepository) CreateDroneMedication(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error) {
	return entity.DroneMedication{}, nil
}

func (cdb MockedDroneActionRepository) Get(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error) {
	return entity.DroneMedication{}, nil
}

func (cdb MockedDroneActionRepository) IsNotFoundErr(err error) bool {
	return true
}

func (cdb MockedDroneActionRepository) GetDroneMedications(ctx context.Context, droneMedication entity.DroneMedication) ([]entity.DroneMedication, error) {
	return []entity.DroneMedication{}, nil
}
