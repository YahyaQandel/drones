package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type MockedDroneNotExistsRepository struct {
	client *gorm.DB
}

func NewMockedDroneNotExistsRepository() repository.IDroneRepo {
	return &MockedDroneNotExistsRepository{}
}

func (cdb MockedDroneNotExistsRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb MockedDroneNotExistsRepository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, gorm.ErrRecordNotFound
}

func (cdb MockedDroneNotExistsRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (cdb MockedDroneNotExistsRepository) Update(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}
func (cdb MockedDroneNotExistsRepository) GetAll(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}
func (cdb MockedDroneNotExistsRepository) GetAvailable(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}

func (cdb MockedDroneNotExistsRepository) GetLoaded(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{}, nil
}
