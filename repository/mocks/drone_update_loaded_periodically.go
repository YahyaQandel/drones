package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type MockedUpdatedLoadedDronePeriodicallyRepository struct {
	drone entity.Drone
}

func NewMockedUpdatedLoadedDronePeriodicallyRepository(receivedDrone entity.Drone) repository.IDroneRepo {
	return &MockedUpdatedLoadedDronePeriodicallyRepository{drone: receivedDrone}
}

func (cdb *MockedUpdatedLoadedDronePeriodicallyRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb *MockedUpdatedLoadedDronePeriodicallyRepository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return cdb.drone, nil
}

func (cdb *MockedUpdatedLoadedDronePeriodicallyRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (cdb *MockedUpdatedLoadedDronePeriodicallyRepository) Update(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	cdb.drone = drone
	return cdb.drone, nil
}

func (cdb MockedUpdatedLoadedDronePeriodicallyRepository) GetAvailable(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}

func (cdb *MockedUpdatedLoadedDronePeriodicallyRepository) GetLoaded(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{
		cdb.drone,
	}, nil
}
