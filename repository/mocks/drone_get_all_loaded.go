package mocks

import (
	"context"

	"drones.com/repository"
	"drones.com/repository/entity"
	usecaseEntity "drones.com/usecase/entity"
	"gorm.io/gorm"
)

type MockedDroneGetAllLoadedRepository struct {
	client *gorm.DB
}

func NewMockedDroneGetAllLoadedRepository() repository.IDroneRepo {
	return &MockedDroneGetAllLoadedRepository{}
}

func (cdb MockedDroneGetAllLoadedRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb MockedDroneGetAllLoadedRepository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{State: string(usecaseEntity.IDLE)}, nil
}
func (cdb MockedDroneGetAllLoadedRepository) GetAll(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}
func (cdb MockedDroneGetAllLoadedRepository) IsNotFoundErr(err error) bool {
	return false
}

func (cdb MockedDroneGetAllLoadedRepository) Update(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}
func (cdb MockedDroneGetAllLoadedRepository) GetAvailable(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}
func (cdb MockedDroneGetAllLoadedRepository) GetLoaded(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{
		{
			ID:              1,
			SerialNumber:    "12345",
			State:           string(usecaseEntity.LOADED),
			Model:           "LightWeight",
			BatteryCapacity: 60,
			Weight:          10,
		},
		{
			ID:              2,
			SerialNumber:    "54321",
			State:           string(usecaseEntity.LOADED),
			Model:           "HeavyWeight",
			BatteryCapacity: 80,
			Weight:          20,
		},
	}, nil
}
