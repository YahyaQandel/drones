package entity

type Medication struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
	Code   string  `json:"code" gorm:"index;uniqueIndex"`
	Image  string  `json:"image"`
}
