package config

import (
	"fmt"
	"log"
	"os"

	"example.com/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// Load environment variable
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		return nil
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbName)

	// Setup from gorm docs
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Checking error
	if err != nil {
		panic("Failed to connect database")
	}

	// Craete table, sync model
	db.AutoMigrate(models.Product{})
	db.AutoMigrate(models.User{})

	return db
}
