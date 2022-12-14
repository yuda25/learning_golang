package controllers

import (
	"learning_golang/initializers"
	"learning_golang/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	// get data of req body
	var body struct {
		Name string
		Stock uint
		Price uint
		CodeProduct string
		CategoryId uint
	}

	c.Bind(&body)

	code := models.CodeProduct{Code: body.CodeProduct}

	// create post
	product := models.Product{Name: body.Name, Stock: body.Stock, Price: body.Price, CodeProduct: code, CategoryId: body.CategoryId, Provider: []models.Provider{
		{Name: "kusuma"},
		{Name: "yuda"},
		{Name: "mubarok"},
	}}

	result := initializers.DB.Preload("Category").Create(&product).Find(&product)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// return it
	c.JSON(200, gin.H{
		"product": product,
	})
}

func GettAllProduct(c *gin.Context)  {
	// get the products
	var products []models.Product
	initializers.DB.Preload("CodeProduct").Preload("Category").Preload("Provider").Find(&products)

	// return it
	c.JSON(200, gin.H{
		"products": products,
	})
}

func GetById(c *gin.Context)  {
	// get id 
	id := c.Param("id")

	// get the products
	var product models.Product
	initializers.DB.Preload("CodeProduct").Preload("Category").Preload("Provider").First(&product, id)

	// return it
	c.JSON(200, gin.H{
		"product": product,
	})
}

func UpdateProduct(c *gin.Context) {
	// get id 
	id := c.Param("id")

	// get data from req body
	var body struct {
		Name string
		Stock uint
		Price uint
	}
	c.Bind(&body)

	// find post were updating
	var product []models.Product
	initializers.DB.Preload("CodeProduct").First(&product, id)

	// update it
	initializers.DB.Model(&product).Updates(models.Product{
		Name: body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	// respond
	c.JSON(200, gin.H{
		"product": product,
	})
}

func DeleteProduct(c *gin.Context) {
	// get id
	id := c.Param("id")
	// detete 
	initializers.DB.Delete(&models.Product{}, id)
	// respond
	c.Status(200)
}