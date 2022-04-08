package main

import (
	"gorm.io/gorm"
)

type DroneMedication struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	DroneSerialNumber string `json:"drone_serial_number" `
	MedicationCode    string `json:"medication_code" gorm:"type:varchar(15)"`
}

// Up is executed when this migration is applied
func Up_20220408065241(txn *gorm.DB) {
	txn.Migrator().CreateTable(DroneMedication{})
}

// Down is executed when this migration is rolled back
func Down_20220408065241(txn *gorm.DB) {
	txn.Migrator().DropTable(DroneMedication{})
}
