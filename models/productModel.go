package models

type Product struct {
	ID uint  `gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
	Stock uint `json:"stock" gorm:"not null"`
	Price uint `json:"price" gorm:"not null"`

	// one to one
	CodeProductId uint `gorm:"ForeignKey:id"`
	CodeProduct CodeProduct

	// one to many
	CategoryId uint `gorm:"ForeignKey:id"`
	Category Category `json:"categories"`

	// many to many
	Provider []Provider `gorm:"many2many:product_providers;"`
}