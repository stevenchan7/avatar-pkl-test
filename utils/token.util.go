package utils

import (
	"log"
	"os"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func generateToken(userID uint) (string, error) {
	// Load env variable
	// path_dir := "C:/Users/Steven Ciam/Documents/avatar-product/utils"
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		return "", err
	}

	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_LIFESPAN"))

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"userID":     userID,
		"exp":        time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}
