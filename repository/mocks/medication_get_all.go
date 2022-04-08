package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type MedicationGetAllRepository struct {
}

func NewMedicationGetAllRepository() repository.IMedicationRepo {
	return &MedicationGetAllRepository{}
}

func (cdb MedicationGetAllRepository) Create(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, nil
}

func (cdb MedicationGetAllRepository) Update(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, nil
}
func (cdb MedicationGetAllRepository) Get(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	if medication.Code == "RX" {
		return entity.Medication{
			ID:     1,
			Name:   "Aspirin",
			Code:   "RX",
			Weight: 100,
			Image:  "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		}, nil
	} else if medication.Code == "UX" {
		return entity.Medication{
			ID:     2,
			Name:   "Advil",
			Code:   "UX",
			Weight: 200,
			Image:  "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		}, nil
	} else if medication.Code == "DX" {
		return entity.Medication{
			ID:     3,
			Name:   "Vicodin",
			Code:   "DX",
			Weight: 300,
			Image:  "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		}, nil
	}
	return entity.Medication{}, nil
}

func (cdb MedicationGetAllRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
