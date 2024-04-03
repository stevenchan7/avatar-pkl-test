package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func VerifyToken(c *gin.Context) {
	bearerToken, err := c.Cookie("Authorization")
	var tokenString string

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "Token missing"})
		c.Abort()
		return
	}

	if len(strings.Split(bearerToken, " ")) == 2 {
		tokenString = strings.Split(bearerToken, " ")[1]
	}

	log.Println(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_JWT")), nil
	})

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "Invalid token"})
		c.Abort()
		return
	}

  // Get token claims and check error
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Attached claims id
		c.Set("userID", claims["userID"])
	} else {
		fmt.Println(err)
		c.Abort()
		return
	}

	c.Next()
}
