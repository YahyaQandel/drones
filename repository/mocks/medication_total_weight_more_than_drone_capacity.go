package mocks

import (
	"context"
	"errors"

	"drones.com/repository"
	"drones.com/repository/entity"
	"gorm.io/gorm"
)

type MedicationTotalWeightMoreThanDroneCapacityRepository struct {
}

func NewMedicationTotalWeightMoreThanDroneCapacityRepository(weight float64) repository.IMedicationRepo {
	return &MedicationTotalWeightMoreThanDroneCapacityRepository{}
}

func (cdb MedicationTotalWeightMoreThanDroneCapacityRepository) Create(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, nil
}

func (cdb MedicationTotalWeightMoreThanDroneCapacityRepository) Update(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return entity.Medication{}, nil
}

func (cdb MedicationTotalWeightMoreThanDroneCapacityRepository) Get(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	if medication.Code == "RX" {
		return entity.Medication{
			ID:     1,
			Name:   "Aspirin",
			Code:   "RX",
			Weight: 20,
			Image:  "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		}, nil
	} else if medication.Code == "PD1" {
		return entity.Medication{
			ID:     1,
			Name:   "Aspirin",
			Code:   "RX",
			Weight: 20,
			Image:  "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		}, nil
	} else if medication.Code == "PD2" {
		return entity.Medication{
			ID:     2,
			Name:   "Advil",
			Code:   "UX",
			Weight: 15,
			Image:  "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		}, nil
	} else if medication.Code == "PD3" {
		return entity.Medication{
			ID:     3,
			Name:   "Vicodin",
			Code:   "DX",
			Weight: 13,
			Image:  "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		}, nil
	}
	return entity.Medication{}, nil
}

func (cdb MedicationTotalWeightMoreThanDroneCapacityRepository) IsNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
