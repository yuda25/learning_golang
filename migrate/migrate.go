package main

import (
	"learning_golang/initializers"
	"learning_golang/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connection()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.Product{},
		&models.CodeProduct{},
	)
}