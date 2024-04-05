package controllers

import (
	"fmt"
	"net/http"

	"example.com/config"
	"example.com/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PassRecoveryHandler(c *gin.Context) {
	db := config.ConnectDB()

	type passRecoveryEmail struct {
		Email string `json:"email" binding:"required"`
	}

	var input passRecoveryEmail

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
	}

	// Populate user find based on email
	var user models.User

	if err := db.First(&user, "email = ?", input.Email).Error; err != nil {
		errMsg := fmt.Sprintf("Failed to find user with email %s: %s", input.Email, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": errMsg})
	}

	verPasswd, _, err := user.NewVerificationPasswd()

	if err != nil {
		errMsg := fmt.Sprintf("Failed to create verification password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": errMsg})
	}

	fmt.Print(user)

	if err := db.Save(&user).Error; err != nil {
		errMsg := fmt.Sprintf("Failed to update user verification hash: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": errMsg})
	}

	subject := "Account recovery"
	HTMLBody :=
		`<html>
            <h1>Click link below to recover password</h1>
            <a href="http://localhost:8080/account-recovery` + user.Username + verPasswd + `">Change password</a>
        <html/>`

	if err := user.SendEmail(subject, HTMLBody); err != nil {
		errMsg := fmt.Sprintf("Failed to send recovery email: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": errMsg})
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Check your email", "verificationPasswd": verPasswd})
}

func ChangePassHandler(c *gin.Context) {
	type ChangePass struct {
		NewPassword1 string `json:"newPassword1"`
		NewPassword2 string `json:"newPassword2"`
		Username     string `json:"username"`
		VerHash      string `json:"verHash"`
	}

	var input ChangePass

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
	}

	db := config.ConnectDB()

	var user models.User

	// Populate user
	if err := db.First(&user, "username = ?", input.Username).Error; err != nil {
		errMsg := fmt.Sprintf("Failed to find user with username %s: %s", input.Username, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": errMsg})
	}

	// Compare incoming request verification password with hashed verification password in db
	if err := bcrypt.CompareHashAndPassword([]byte(user.VerHash), []byte(input.VerHash)); err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
		errMsg := fmt.Sprintf("Invalid verification password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": errMsg})
	}

	// Check if password1 and 2 is match
	if input.NewPassword1 != input.NewPassword2 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Both password must match"})
	}

	// Hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword1), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "There was an issue reseting password"})
	}

	// Update user password in db
	user.Password = string(hashedPassword)

	if err := db.Save(&user).Error; err != nil {
		errMsg := fmt.Sprintf("Failed to update user password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": errMsg})
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Successfully changed password"})
}
