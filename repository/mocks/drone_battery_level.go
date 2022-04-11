package mocks

import (
	"context"

	"drones.com/repository"
	"drones.com/repository/entity"
	usecaseEntity "drones.com/usecase/entity"
	"gorm.io/gorm"
)

type MockedDroneBatteryLevelRepository struct {
	client *gorm.DB
}

func NewMockedDroneBatteryLevelRepository() repository.IDroneRepo {
	return &MockedDroneBatteryLevelRepository{}
}

func (cdb MockedDroneBatteryLevelRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb MockedDroneBatteryLevelRepository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{
		ID:              1,
		SerialNumber:    "12345",
		State:           string(usecaseEntity.IDLE),
		Model:           "LightWeight",
		BatteryCapacity: 60,
		Weight:          10,
	}, nil
}

func (cdb MockedDroneBatteryLevelRepository) IsNotFoundErr(err error) bool {
	return false
}

func (cdb MockedDroneBatteryLevelRepository) Update(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb MockedDroneBatteryLevelRepository) GetAvailable(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}
