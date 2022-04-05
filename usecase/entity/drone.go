package entity

type Drone struct {
	SerialNumber    string `json:"serial_number" valid:"required,alpha,maxstringlength(100)~serial_number: maximum length is 100"`
	Model           string
	Weight          float64 `json:"weight" valid:"required,range(1|500)"`
	BatteryCapacity float64
	State           string
}

var DroneModels = []string{"Lightweight", "Middleweight", "Cruiserweight", "Heavyweight"}
