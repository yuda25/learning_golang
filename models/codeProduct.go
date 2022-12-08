package models


type CodeProduct struct {
	ID uint `gorm:"primary_key"`
	Code string `json:"code" gorm:"not null"`
}