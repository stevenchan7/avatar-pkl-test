package utils

import (
	"example.com/config"
	"example.com/models"
	"golang.org/x/crypto/bcrypt"
)

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username, password string) (string, error) {
	DB := config.ConnectDB()
	var u models.User

	// Find user by username
	DB.First(&u, "username = ?", username)

	// Verify password
	if err := verifyPassword(password, u.Password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	tokenString, err := generateToken(u.ID)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
