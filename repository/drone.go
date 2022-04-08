package repository

import (
	"context"
	"errors"

	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type IDroneRepo interface {
	Create(ctx context.Context, drone entity.Drone) (entity.Drone, error)
	Get(ctx context.Context, drone entity.Drone) (entity.Drone, error)
	Update(ctx context.Context, drone entity.Drone) (entity.Drone, error)
	IsNotFoundErr(err error) bool
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

func (cdb DroneRepository) Update(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	cdb.client.WithContext(ctx).Model(&drone).Updates(drone)
	return drone, nil
}

func (cdb DroneRepository) Get(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	droneResponse := entity.Drone{}
	result := cdb.client.WithContext(ctx).Where(&entity.Drone{SerialNumber: drone.SerialNumber}).Last(&droneResponse)
	if cdb.IsNotFoundErr(result.Error) {
		return entity.Drone{}, result.Error
	}
	if result.Error != nil {
		return entity.Drone{}, result.Error
	}
	return droneResponse, nil
}

func (cdb DroneRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
