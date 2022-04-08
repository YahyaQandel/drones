package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type MedicationOverWeightRepository struct {
	weight float64
}

func NewMedicationOverWeightRepository(weight float64) repository.IMedicationRepo {
	return &MedicationOverWeightRepository{weight: weight}
}

func (cdb MedicationOverWeightRepository) Create(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, nil
}

func (cdb MedicationOverWeightRepository) Update(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, nil
}

func (cdb MedicationOverWeightRepository) Get(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return entity.Medication{Weight: cdb.weight}, nil
}

func (cdb MedicationOverWeightRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
