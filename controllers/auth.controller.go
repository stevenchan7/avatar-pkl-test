package controllers

import (
	"net/http"

	"example.com/config"
	"example.com/models"
	"example.com/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	DB := config.ConnectDB()
	var userBinder models.RegisterUserInput

	if err := c.ShouldBindJSON(&userBinder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "err": err.Error()})
		return
	}

	newUser := models.User{Username: userBinder.Username, Password: userBinder.Password}

	if err := DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": newUser})
}

func Login(c *gin.Context) {
	var userBinder models.LoginUserInput

	if err := c.ShouldBind(&userBinder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "err": err})
    return
	}

	token, err := utils.LoginCheck(userBinder.Username, userBinder.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "err": err})
		return
	}

	// Set token within cookie
	bearerToken := "Bearer " + token

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", bearerToken, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"success": true, "token": token})
}

func Logout(c *gin.Context) {
	_, err := c.Cookie("Authorization")

	if err == nil {
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Successfully removed token"})
	}
}
