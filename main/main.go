package main

import (
	"log"
	"net/http"
	"os"

	"example.com/controllers"
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// Load environment variable
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
			return
		}

		secret := os.Getenv("SECRET_JWT")
		c.JSON(http.StatusOK, gin.H{"message": "Hello world", "Secret": secret})
	})

	r.GET("/todo-item", controllers.GetTodoItem)

	r.GET("/todo-item/:id", controllers.GetTodoItemByID)

	admin := r.Group("/admin")
	admin.Use(middlewares.VerifyToken)
	{
		admin.GET("/aaa", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Hi from protected"})
		})

		admin.POST("/add-todo-item", controllers.PostTodoItem)

		admin.POST("/update-todo-item", controllers.UpdateTodoItem)

		admin.POST("/update-todo-item-status", controllers.UpdateTodoItemStatus)

		admin.POST("/delete-todo-item", controllers.DeleteToDoItem)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)

		auth.POST("/login", controllers.Login)

		auth.POST("/logout", controllers.Logout)
	}

	r.POST("/pass-recovery", controllers.PassRecoveryHandler)

	r.POST("/change-pass", controllers.ChangePassHandler)

	r.Run()
}
