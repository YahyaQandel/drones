package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type DroneNotExistsRepository struct {
	client *gorm.DB
}

func NewMockedDroneNotExistsRepository() repository.IDroneRepo {
	return &DroneNotExistsRepository{}
}

func (cdb DroneNotExistsRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (cdb DroneNotExistsRepository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return entity.Drone{}, gorm.ErrRecordNotFound
}

func (cdb DroneNotExistsRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
