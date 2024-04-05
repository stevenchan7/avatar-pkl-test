package controllers

import (
	"fmt"
	"net/http"

	"example.com/config"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

func PostTodoItem(c *gin.Context) {
	// Call ConnectDB
	DB := config.ConnectDB()

	var input models.AddTodoItemInput

	// Bind input and check error at the same time
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
	}

	newTodoItem := models.TodoItem{Title: input.Title, Desc: input.Desc, Status: input.Status, DueDate: input.DueDate}

	// Insert new product to database
	if err := DB.Create(&newTodoItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": newTodoItem})
}

func UpdateTodoItem(c *gin.Context) {
	DB := config.ConnectDB()
	// Bind user input
	var input models.UpdateTodoItemInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "err": err.Error()})
	}

	// Find product in database
	TodoItemId := input.ID

	var updatedTodoItem models.TodoItem

	DB.First(&updatedTodoItem, "id = ?", TodoItemId)

	// Update product in database
	// updatedProduct.UpdatedAt = time.Now()
	DB.Model(&updatedTodoItem).Updates(input)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": updatedTodoItem})
}

func UpdateTodoItemStatus(c *gin.Context) {
	DB := config.ConnectDB()
	// Bind user input
	var input models.UpdateTodoItemStatusInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "err": err.Error()})
	}

	// Find product in database
	TodoItemId := input.ID

	var updatedTodoItem models.TodoItem

	DB.First(&updatedTodoItem, "id = ?", TodoItemId)

	// Update product in database
	// updatedProduct.UpdatedAt = time.Now()
	DB.Model(&updatedTodoItem).Updates(input)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": updatedTodoItem})
}

func DeleteToDoItem(c *gin.Context) {
	DB := config.ConnectDB()

	var input models.DeleteTodoItemInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "err": err.Error()})
	}

	// Delete from database
	DB.Delete(&models.TodoItem{}, input.ID)

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": fmt.Sprintf("Successfully delete prduct with ID %s", input.ID)})
}
