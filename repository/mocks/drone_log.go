package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type MockedDroneLogRepository struct {
	client *gorm.DB
}

func NewMockedDroneLogRepository() repository.IDroneLogRepo {
	return &MockedDroneLogRepository{}
}

func (cdb MockedDroneLogRepository) Create(ctx context.Context, droneLog entity.DroneLog) (entity.DroneLog, error) {
	return entity.DroneLog{}, nil
}

func (cdb MockedDroneLogRepository) Get(ctx context.Context, droneLog entity.DroneLog) (entity.DroneLog, error) {
	return entity.DroneLog{}, nil
}

func (cdb MockedDroneLogRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
