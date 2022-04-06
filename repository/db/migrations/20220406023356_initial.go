package main

import (
	"gorm.io/gorm"
)

type Drone struct {
	ID              uint    `json:"-" gorm:"primaryKey"`
	SerialNumber    string  `json:"serial_number" gorm:"index;uniqueIndex"`
	Model           string  `json:"model" gorm:"type:varchar(15)"`
	Weight          float64 `json:"weight"`
	BatteryCapacity float64
	State           string `json:"state" gorm:"default:IDLE"`
}

// Up is executed when this migration is applied
func Up_20220406023356(txn *gorm.DB) {

	txn.Migrator().CreateTable(Drone{})
}

// Down is executed when this migration is rolled back
func Down_20220406023356(txn *gorm.DB) {
	txn.Migrator().DropTable(Drone{})

}
