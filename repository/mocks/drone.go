package mocks

import (
	"context"

	"drones.com/repository"
	"drones.com/repository/entity"
	usecaseEntity "drones.com/usecase/entity"
	"gorm.io/gorm"
)

type DroneRepository struct {
	client *gorm.DB
}

func NewMockedDroneRepository() repository.IDroneRepo {
	return &DroneRepository{}
}

func (cdb DroneRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb DroneRepository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{State: string(usecaseEntity.IDLE)}, nil
}
func (cdb DroneRepository) GetAll(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}
func (cdb DroneRepository) IsNotFoundErr(err error) bool {
	return false
}

func (cdb DroneRepository) Update(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb DroneRepository) GetAvailable(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}

func (cdb DroneRepository) GetLoaded(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}
