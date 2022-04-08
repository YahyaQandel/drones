package repository

import (
	"context"

	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type IMedicationRepo interface {
	Create(ctx context.Context, medication entity.Medication) (entity.Medication, error)
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
