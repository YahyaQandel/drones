package entity

type Drone struct {
	SerialNumber    string  `json:"serial_number" valid:"required,alpha,maxstringlength(100)~serial_number: maximum length is 100"`
	Model           string  `json:"model"`
	Weight          float64 `json:"weight" valid:"required,range(1|500)"`
	BatteryCapacity float64 `json:"battery_capacity"`
	State           string  `json:"state"`
}
type modelType string
type stateType string

const (
	Lightweight   modelType = "Lightweight"
	Middleweight  modelType = "Middleweight"
	Cruiserweight modelType = "Cruiserweight"
	Heavyweight   modelType = "Heavyweight"
)

const (
	IDLE       stateType = "IDLE"
	LOADING    stateType = "LOADING"
	LOADED     stateType = "LOADED"
	DELIVERING stateType = "DELIVERING"
	DELIVERED  stateType = "DELIVERED"
	RETURNING  stateType = "RETURNING"
)

var DroneModels = []modelType{Lightweight, Middleweight, Cruiserweight, Heavyweight}
var DroneStates = []stateType{IDLE, LOADING, LOADED, DELIVERING, DELIVERED, RETURNING}
