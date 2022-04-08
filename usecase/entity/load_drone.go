package entity

type LoadDrone struct {
	DroneSerialNumber string `json:"drone_serial_number"  valid:"required"`
	MedicationCode    string `json:"medication_code" valid:"required"`
}
