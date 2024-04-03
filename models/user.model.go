package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RegisterUserInput struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
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
