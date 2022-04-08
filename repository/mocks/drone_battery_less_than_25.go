package mocks

import (
	"context"

	"drones.com/repository"
	"drones.com/repository/entity"
)

type DroneBatteryLessThan25Repository struct {
	batteryCapacity float64
}

func NewDroneBatteryLessThan25Repository(batteryCapacity float64) repository.IDroneRepo {
	return &DroneBatteryLessThan25Repository{batteryCapacity: batteryCapacity}
}

func (cdb DroneBatteryLessThan25Repository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb DroneBatteryLessThan25Repository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{BatteryCapacity: cdb.batteryCapacity}, nil
}

func (cdb DroneBatteryLessThan25Repository) IsNotFoundErr(err error) bool {
	return false
}
