package mocks

import (
	"context"

	"drones.com/repository"
	"drones.com/repository/entity"
	usecaseEntity "drones.com/usecase/entity"
)

type MockedDroneLessWeightThanMedicationRepository struct {
	weight float64
}

func NewMockedDroneLessWeightThanMedicationRepository(weight float64) repository.IDroneRepo {
	return &MockedDroneLessWeightThanMedicationRepository{weight: weight}
}

func (cdb MockedDroneLessWeightThanMedicationRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb MockedDroneLessWeightThanMedicationRepository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{Weight: cdb.weight, State: string(usecaseEntity.IDLE)}, nil
}

func (cdb MockedDroneLessWeightThanMedicationRepository) IsNotFoundErr(err error) bool {
	return false
}

func (cdb MockedDroneLessWeightThanMedicationRepository) Update(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}
