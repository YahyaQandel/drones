package mocks

import (
	"context"

	"drones.com/repository"
	"drones.com/repository/entity"
	usecaseEntity "drones.com/usecase/entity"
	"gorm.io/gorm"
)

type MockedDroneGetAllAvailableRepository struct {
	client *gorm.DB
}

func NewMockedDroneGetAllAvailableRepository() repository.IDroneRepo {
	return &MockedDroneGetAllAvailableRepository{}
}

func (cdb MockedDroneGetAllAvailableRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb MockedDroneGetAllAvailableRepository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{State: string(usecaseEntity.IDLE)}, nil
}
func (cdb MockedDroneGetAllAvailableRepository) GetAll(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}
func (cdb MockedDroneGetAllAvailableRepository) IsNotFoundErr(err error) bool {
	return false
}

func (cdb MockedDroneGetAllAvailableRepository) Update(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb MockedDroneGetAllAvailableRepository) GetAvailable(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{
		{
			ID:              1,
			SerialNumber:    "12345",
			State:           string(usecaseEntity.IDLE),
			Model:           "LightWeight",
			BatteryCapacity: 60,
			Weight:          10,
		},
		{
			ID:              2,
			SerialNumber:    "54321",
			State:           string(usecaseEntity.IDLE),
			Model:           "HeavyWeight",
			BatteryCapacity: 80,
			Weight:          20,
		},
	}, nil
}

func (cdb MockedDroneGetAllAvailableRepository) GetLoaded(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}
