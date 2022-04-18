package main

import (
	"gorm.io/gorm"
)

// Up is executed when this migration is applied
func Up_20220416063731(txn *gorm.DB) {
	type DroneLog struct {
		ID                uint    `json:"id" gorm:"primaryKey"`
		DroneSerialNumber string  `json:"drone_serial_number"`
		DroneBatteryLevel float64 `json:"drone_battery_level"`
		DroneState        string  `json:"drone_state"`
	}

	txn.Migrator().CreateTable(DroneLog{})
}

// Down is executed when this migration is rolled back
func Down_20220416063731(txn *gorm.DB) {
	type DroneLog struct {
		ID                uint    `json:"id" gorm:"primaryKey"`
		DroneSerialNumber string  `json:"drone_serial_number"`
		DroneBatteryLevel float64 `json:"drone_battery_level"`
		DroneState        string  `json:"drone_state"`
	}

	txn.Migrator().DropTable(DroneLog{})
}
