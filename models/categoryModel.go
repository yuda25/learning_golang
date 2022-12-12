package models

type Category struct {
	ID uint `gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`

	Product []Product
}