package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	usecaseEntity "drones.com/usecase/entity"
	"gorm.io/gorm"
)

type MockedLoadedDroneRepository struct {
	drone entity.Drone
}

func NewMockedLoadedDroneRepository() repository.IDroneRepo {
	return &MockedDroneExistsRepository{drone: entity.Drone{SerialNumber: "XDX", BatteryCapacity: 30, Weight: 100, State: string(usecaseEntity.LOADED)}}
}

func (cdb MockedLoadedDroneRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb MockedLoadedDroneRepository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{State: string(usecaseEntity.LOADED)}, nil
}

func (cdb MockedLoadedDroneRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (cdb MockedLoadedDroneRepository) Update(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb MockedLoadedDroneRepository) GetAvailable(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}
