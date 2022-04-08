package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type MedicationNotExistsRepository struct {
}

func NewMockedMedicationNotExistsRepository() repository.IMedicationRepo {
	return &MedicationNotExistsRepository{}
}

func (cdb MedicationNotExistsRepository) Create(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, nil
}

func (cdb MedicationNotExistsRepository) Update(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, nil
}

func (cdb MedicationNotExistsRepository) Get(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, gorm.ErrRecordNotFound
}

func (cdb MedicationNotExistsRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
