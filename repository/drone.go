package repository

import (
	"context"

	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type IDroneRepo interface {
	Create(ctx context.Context, drone entity.Drone) (entity.Drone, error)
}

type DroneRepository struct {
	client *gorm.DB
}

func NewDroneRepository(client *gorm.DB) IDroneRepo {
	return &DroneRepository{client: client}
}

func (cdb DroneRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	result := cdb.client.WithContext(ctx).Create(&drone)
	if result.Error != nil {
		return entity.Drone{}, result.Error
	}
	return drone, nil
}
