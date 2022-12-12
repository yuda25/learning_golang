package models

type Provider struct {
	ID uint  `gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
}