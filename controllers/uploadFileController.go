package controllers

import (
	"fmt"
	"learning_golang/initializers"
	"learning_golang/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)


func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	filename := filepath.Base(file.Filename)
	path := "././uploads/" + filename
	fmt.Println(path)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	upload := models.Upload{Path: path}

	result := initializers.DB.Create(&upload)

	if result.Error != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	c.String(http.StatusOK, "'%s' uploaded!", file.Filename)
}