package controllers

import (
	"fmt"
	"io"
	"learning_golang/initializers"
	"learning_golang/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

func CreateExcel(c *gin.Context) {
	f := excelize.NewFile()
    // Create a new sheet.
    index := f.NewSheet("Sheet1")
    // Set value of a cell.
	var products []models.Product
	initializers.DB.Find(&products)
	for i, data := range products {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+1), data.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+1), data.Price)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+1), data.Stock)
	}

    // Set active sheet of the workbook.
    f.SetActiveSheet(index)
    // Save spreadsheet by the given path.
	id := uuid.New()
	path := "assets/product"+id.String()+".xlsx"
	fmt.Println(path)

    if err := f.SaveAs(path); err != nil {
        fmt.Println(err)
    }

	// insert ke db
	excel := models.Excel{Path: path}
	result := initializers.DB.Create(&excel)

	if result.Error != nil {
		c.String(http.StatusInternalServerError, "Error insert to db")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "data excel created",
		"path": excel.Path,
	})
}

func DownloadExcel(c *gin.Context) {
	// cari data dari db
	id := c.Param("id")
	var excel models.Excel
	initializers.DB.First(&excel, id)

	// Buka file Excel yang akan didownload
	file, err := os.Open(excel.Path)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error opening file")
		return
	}
	defer file.Close()

	// Set Content-Type header sebagai application/vnd.ms-excel untuk menandakan bahwa
	// file yang akan didownload adalah file Excel
	c.Header("Content-Type", "application/vnd.ms-excel")

	// Set Content-Disposition header dengan attachment dan nama file sehingga file
	// akan otomatis didownload ketika API ini dipanggil
	c.Header("Content-Disposition", "attachment; filename=product.xlsx")

	// Menuliskan isi file ke response body sehingga file dapat didownload
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error writing file")
		return
	}
}
