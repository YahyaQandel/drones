package entity

type Drone struct {
	ID              uint    `json:"id" gorm:"primaryKey"`
	SerialNumber    string  `json:"serial_number" gorm:"index;uniqueIndex"`
	Model           string  `json:"model" gorm:"type:varchar(15)"`
	Weight          float64 `json:"weight"`
	BatteryCapacity float64 `json:"battery_capacity"`
	State           string  `json:"state" gorm:"default:IDLE"`
}
