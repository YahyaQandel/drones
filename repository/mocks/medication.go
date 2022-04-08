package mocks

import (
	"context"

	"drones.com/repository"
	"drones.com/repository/entity"
)

type MedicationRepository struct {
}

func NewMockedMedicationRepository() repository.IMedicationRepo {
	return &MedicationRepository{}
}

func (cdb MedicationRepository) Create(ctx context.Context, drone entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, nil
}
