package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	usecaseEntity "drones.com/usecase/entity"
	"gorm.io/gorm"
)

type MockedLoadedDroneExistsRepository struct {
	drone entity.Drone
}

func NewMockedLoadedDroneExistsRepository() repository.IDroneRepo {
	return &MockedLoadedDroneExistsRepository{drone: entity.Drone{SerialNumber: "XDX", BatteryCapacity: 30, Weight: 100, State: string(usecaseEntity.LOADED)}}
}

func (cdb MockedLoadedDroneExistsRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb MockedLoadedDroneExistsRepository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return cdb.drone, nil
}

func (cdb MockedLoadedDroneExistsRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (cdb MockedLoadedDroneExistsRepository) Update(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	cdb.drone.State = string(usecaseEntity.LOADED)
	return cdb.drone, nil
}
