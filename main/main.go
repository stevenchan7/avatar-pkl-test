package main

import (
	"net/http"
	"os"

	"example.com/controllers"
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		secret := os.Getenv("TOKEN_SECRET")
		c.JSON(http.StatusOK, gin.H{"message": "Hello sayang", "Secret": secret})
	})

	r.GET("/products", controllers.GetProducts)
  
  r.GET("/products/:prodID", controllers.GetProductByID)

	admin := r.Group("/admin")
	admin.Use(middlewares.VerifyToken)
	{
		admin.POST("/add-product", controllers.PostProduct)

		admin.POST("/edit-product", controllers.EditProduct)

    admin.POST("/delete-product", controllers.DeleteProduct)

		admin.GET("/protected", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Hi from protected", "user_id": userID})
		})
	}

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)

		auth.POST("/login", controllers.Login)

		auth.POST("/logout", controllers.Logout)
	}

	r.Run()
}
