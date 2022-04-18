package repository

import (
	"context"
	"errors"

	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type IDroneLogRepo interface {
	Create(ctx context.Context, droneLog entity.DroneLog) (entity.DroneLog, error)
	Get(ctx context.Context, droneLog entity.DroneLog) (entity.DroneLog, error)
	IsNotFoundErr(err error) bool
}

type DroneLogRepository struct {
	client *gorm.DB
}

func NewDroneLogRepository(client *gorm.DB) IDroneLogRepo {
	return &DroneLogRepository{client: client}
}

func (cdb DroneLogRepository) Create(ctx context.Context, droneLog entity.DroneLog) (entity.DroneLog, error) {
	result := cdb.client.WithContext(ctx).Create(&droneLog)
	if result.Error != nil {
		return entity.DroneLog{}, result.Error
	}
	return droneLog, nil
}

// TODO: refactor to GetLogBySerialNumber
func (cdb DroneLogRepository) Get(ctx context.Context, droneLog entity.DroneLog) (entity.DroneLog, error) {
	droneResponse := entity.DroneLog{}
	result := cdb.client.WithContext(ctx).Where(&entity.DroneLog{DroneSerialNumber: droneLog.DroneSerialNumber}).Last(&droneResponse)
	if cdb.IsNotFoundErr(result.Error) {
		return entity.DroneLog{}, result.Error
	}
	if result.Error != nil {
		return entity.DroneLog{}, result.Error
	}
	return droneResponse, nil
}

func (cdb DroneLogRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
