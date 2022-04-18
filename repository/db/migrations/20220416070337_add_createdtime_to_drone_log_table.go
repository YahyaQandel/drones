package main

import (
	"time"

	"gorm.io/gorm"
)

// Up is executed when this migration is applied
func Up_20220416070337(txn *gorm.DB) {
	type DroneLog struct {
		CreatedTime time.Time
	}
	txn.Migrator().AddColumn(DroneLog{}, "created_time")
}

// Down is executed when this migration is rolled back
func Down_20220416070337(txn *gorm.DB) {
	type DroneLog struct {
		CreatedTime time.Time
	}
	txn.Migrator().DropColumn(DroneLog{}, "created_time")
}
