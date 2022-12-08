package models

type Product struct {
	ID uint  `gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
	Stock uint `json:"stock" gorm:"not null"`
	Price uint `json:"price" gorm:"not null"`

	CodeProductId uint `gorm:"ForeignKey:id"`
	CodeProduct CodeProduct
}