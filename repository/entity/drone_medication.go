package entity

type DroneMedication struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	DroneSerialNumber string `json:"drone_serial_number" gorm:"index;uniqueIndex"`
	MedicationCode    string `json:"medication_code" gorm:"type:varchar(15)"`
}
