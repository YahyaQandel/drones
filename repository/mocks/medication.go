package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type MedicationRepository struct {
}

func NewMockedMedicationRepository() repository.IMedicationRepo {
	return &MedicationRepository{}
}

func (cdb MedicationRepository) Create(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, nil
}

func (cdb MedicationRepository) Update(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, nil
}
func (cdb MedicationRepository) Get(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, nil
}

func (cdb MedicationRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
