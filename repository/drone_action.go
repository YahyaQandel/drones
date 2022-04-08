package repository

import (
	"context"
	"errors"

	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type IDroneActionRepo interface {
	CreateDroneMedication(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error)
	Get(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error)
	IsNotFoundErr(err error) bool
}

type DroneActionRepository struct {
	client *gorm.DB
}

func NewDroneActionRepository(client *gorm.DB) IDroneActionRepo {
	return &DroneActionRepository{client: client}
}

func (cdb DroneActionRepository) CreateDroneMedication(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error) {
	result := cdb.client.WithContext(ctx).Create(&droneMedication)
	if result.Error != nil {
		return entity.DroneMedication{}, result.Error
	}
	return droneMedication, nil
}

func (cdb DroneActionRepository) Get(ctx context.Context, droneMedication entity.DroneMedication) (entity.DroneMedication, error) {
	droneMedicationResponse := entity.DroneMedication{}
	result := cdb.client.WithContext(ctx).Where(&entity.DroneMedication{DroneSerialNumber: droneMedication.DroneSerialNumber}).Last(&droneMedicationResponse)
	if cdb.IsNotFoundErr(result.Error) {
		return entity.DroneMedication{}, result.Error
	}
	if result.Error != nil {
		return entity.DroneMedication{}, result.Error
	}
	return droneMedicationResponse, nil
}

func (cdb DroneActionRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
