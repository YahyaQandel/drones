package repository

import (
	"context"
	"errors"

	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type IMedicationRepo interface {
	Create(ctx context.Context, medication entity.Medication) (entity.Medication, error)
	Update(ctx context.Context, medication entity.Medication) (entity.Medication, error)
	Get(ctx context.Context, medication entity.Medication) (entity.Medication, error)
	IsNotFoundErr(err error) bool
}

type MedicationRepository struct {
	client *gorm.DB
}

func NewMedicationRepository(client *gorm.DB) IMedicationRepo {
	return &MedicationRepository{client: client}
}

func (cdb MedicationRepository) Create(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	result := cdb.client.WithContext(ctx).Create(&medication)
	if result.Error != nil {
		return entity.Medication{}, result.Error
	}
	return medication, nil
}

func (cdb MedicationRepository) Update(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	result := cdb.client.WithContext(ctx).Save(&medication)
	if result.Error != nil {
		return entity.Medication{}, result.Error
	}
	return medication, nil
}

func (cdb MedicationRepository) Get(ctx context.Context, drone entity.Medication) (entity.Medication, error) {
	medicationResponse := entity.Medication{}
	result := cdb.client.WithContext(ctx).Where(&entity.Medication{Code: drone.Code}).Last(&medicationResponse)
	if cdb.IsNotFoundErr(result.Error) {
		return entity.Medication{}, result.Error
	}
	if result.Error != nil {
		return entity.Medication{}, result.Error
	}
	return medicationResponse, nil
}

func (cdb MedicationRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
