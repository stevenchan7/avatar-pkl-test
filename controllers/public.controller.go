package controllers

import (
	"net/http"

	"example.com/config"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

func GetTodoItem(c *gin.Context) {
	DB := config.ConnectDB()
	var todoItems []models.TodoItem

	if err := DB.Select("title", "desc", "status", "due_date").Find(&todoItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": todoItems})
}

func GetTodoItemByID(c *gin.Context) {
	DB := config.ConnectDB()

	// Get ID in URL parameter
	id := c.Param("id")

	var TodoItem models.TodoItem

	// Check error
	if err := DB.Select("title", "desc", "status", "due_date").First(&TodoItem, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": "No Todo item with this ID"})
		return
	}

	// Send product
	c.JSON(http.StatusOK, gin.H{"success": true, "data": TodoItem})
}
