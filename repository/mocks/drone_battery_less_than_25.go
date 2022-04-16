package mocks

import (
	"context"

	"drones.com/repository"
	"drones.com/repository/entity"
	usecaseEntity "drones.com/usecase/entity"
)

type MockedDroneBatteryLessThan25Repository struct {
	batteryCapacity float64
}

func NewDroneBatteryLessThan25Repository(batteryCapacity float64) repository.IDroneRepo {
	return &MockedDroneBatteryLessThan25Repository{batteryCapacity: batteryCapacity}
}

func (cdb MockedDroneBatteryLessThan25Repository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb MockedDroneBatteryLessThan25Repository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{BatteryCapacity: cdb.batteryCapacity, State: string(usecaseEntity.IDLE)}, nil
}

func (cdb MockedDroneBatteryLessThan25Repository) IsNotFoundErr(err error) bool {
	return false
}

func (cdb MockedDroneBatteryLessThan25Repository) Update(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb MockedDroneBatteryLessThan25Repository) GetAvailable(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}

func (cdb MockedDroneBatteryLessThan25Repository) GetLoaded(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}
