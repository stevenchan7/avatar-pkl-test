package controllers

import (
	"net/http"

	"example.com/config"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	DB := config.ConnectDB()
	var products []models.Product

	if err := DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": products})
}

func GetProductByID(c *gin.Context) {
	DB := config.ConnectDB()

	// Get ID in URL parameter
	prodID := c.Param("prodID")

	var product models.Product

  // Check error
	if err := DB.First(&product, "id = ?", prodID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": "No product with this ID"})
		return
	}

	// Send product
	c.JSON(http.StatusOK, gin.H{"success": true, "data": product})
}
