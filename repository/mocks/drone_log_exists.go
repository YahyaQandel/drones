package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	usecaseEntity "drones.com/usecase/entity"
	"gorm.io/gorm"
)

type MockedDroneLogExistsRepository struct {
	droneLog entity.DroneLog
}

func NewMockedDroneLogExistsRepository() repository.IDroneLogRepo {
	return &MockedDroneLogExistsRepository{}
}

func (cdb *MockedDroneLogExistsRepository) Create(ctx context.Context, droneLog entity.DroneLog) (entity.DroneLog, error) {
	cdb.droneLog = entity.DroneLog{
		DroneSerialNumber: "XDX",
		DroneBatteryLevel: 30,
		DroneState:        string(usecaseEntity.LOADED),
	}
	return cdb.droneLog, nil
}

func (cdb *MockedDroneLogExistsRepository) Get(ctx context.Context, droneLog entity.DroneLog) (entity.DroneLog, error) {
	return cdb.droneLog, nil
}

func (cdb MockedDroneLogExistsRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
