package models

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	VerHash   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RegisterUserInput struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
}

type LoginUserInput struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// replace plain password with hashed password
	u.Password = string(hashedPassword)

	return nil
}

// Make random verification string (plain and hashed)
func (u *User) NewVerificationPasswd() (verPasswd string, verHash string, err error) {
	alphaNumRune := []rune("ajsdSJDAbsadASDBSAJKsd8296473146")
	verRandRune := make([]rune, 64)

	for i := 0; i < 64; i++ {
		verRandRune[i] = alphaNumRune[rand.Intn(len(alphaNumRune)-1)]
	}

	fmt.Println(verRandRune)

	verPassword := string(verRandRune)
	fmt.Println(verPassword)

	verPasswordHash, err := bcrypt.GenerateFromPassword([]byte(verPassword), bcrypt.DefaultCost)

	if err != nil {
		return verPassword, string(verPasswordHash), err
	}

	// Store hashed verification password to database
	u.VerHash = string(verPasswordHash)

	return verPassword, string(verPasswordHash), nil
}

func (u *User) SendEmail(subject string, HTMLBody string) error {
	// Load environment variable
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		return nil
	}

	auth := smtp.PlainAuth("", "mismag12345@gmail.com", os.Getenv("APP_KEY"), "smtp.gmail.com")

	to := []string{u.Email}

	msg := []byte(
		"To: " + u.Email + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME: MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
			HTMLBody)

	err := smtp.SendMail("smtp.gmail.com:587", auth, "mismag12345@gmail.com", to, msg)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
