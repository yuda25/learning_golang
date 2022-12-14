package main

import (
	"learning_golang/controllers"
	"learning_golang/initializers"
	"learning_golang/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connection()
}

func main() {
	r := gin.Default()
	r.POST("/create-product", middleware.RequireAuth, controllers.CreateProduct)
	r.GET("/get-all", middleware.RequireAuth, controllers.GettAllProduct)
	r.GET("/get-one/:id", middleware.RequireAuth, controllers.GetById)
	r.PUT("/update-product/:id", middleware.RequireAuth, controllers.UpdateProduct)
	r.DELETE("/delete-product/:id", middleware.RequireAuth, controllers.DeleteProduct)

	r.POST("/auth/signup", controllers.SignUp)
	r.POST("/auth/signin", controllers.SignIn)
	r.GET("/auth/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
