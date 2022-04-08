package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type DroneExistsRepository struct {
}

func NewMockedDroneExistsRepository() repository.IDroneRepo {
	return &DroneExistsRepository{}
}

func (cdb DroneExistsRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb DroneExistsRepository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{SerialNumber: "XDX", BatteryCapacity: 30, Weight: 100}, nil
}

func (cdb DroneExistsRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
