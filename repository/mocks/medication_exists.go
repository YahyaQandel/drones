package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type MedicationExistsRepository struct {
}

func NewMedicationExistsRepository() repository.IMedicationRepo {
	return &MedicationExistsRepository{}
}

func (cdb MedicationExistsRepository) Create(ctx context.Context, drone entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, nil
}

func (cdb MedicationExistsRepository) Get(ctx context.Context, drone entity.Medication) (entity.Medication, error) {
	return entity.Medication{Code: "RX", Weight: 10}, nil
}

func (cdb MedicationExistsRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
