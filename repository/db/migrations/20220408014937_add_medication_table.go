package main

import (
	"gorm.io/gorm"
)

type Medication struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
	Code   string  `json:"code" gorm:"index;uniqueIndex"`
	Image  string  `json:"image"`
}

// Up is executed when this migration is applied
func Up_20220408014937(txn *gorm.DB) {
	txn.Migrator().CreateTable(Medication{})
}

// Down is executed when this migration is rolled back
func Down_20220408014937(txn *gorm.DB) {
	txn.Migrator().DropTable(Medication{})
}
