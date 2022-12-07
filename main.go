package main

import (
	"learning_golang/controllers"
	"learning_golang/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connection()
}

func main() {
	r := gin.Default()
	r.POST("/create-product", controllers.CreateProduct)
	r.GET("/get-all", controllers.GettAllProduct)
	r.GET("/get-one/:id", controllers.GetById)
	r.PUT("/update-product/:id", controllers.UpdateProduct)
	r.DELETE("/delete-product/:id", controllers.DeleteProduct)
	r.Run()
}
