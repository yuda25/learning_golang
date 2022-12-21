package models

type Excel struct {
	ID uint  `gorm:"primary_key"`
	Path string `json:"path" gorm:"not null"`
}