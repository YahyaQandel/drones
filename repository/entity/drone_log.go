package entity

import "time"

type DroneLog struct {
	ID                uint    `json:"id" gorm:"primaryKey"`
	DroneSerialNumber string  `json:"drone_serial_number"`
	DroneBatteryLevel float64 `json:"drone_battery_level"`
	DroneState        string  `json:"drone_state"`
	CreatedTime       time.Time
}
