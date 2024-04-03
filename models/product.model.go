package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Image     string `json:"image"`
	Playstore string `json:"playstore"`
	Appstote  string `json:"appstore"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ProductInput struct {
	Title     string `json:"title" form:"title"`
	Desc      string `json:"desc" form:"desc"`
	Image     string `json:"image" form:"image"`
	Playstore string `json:"playstore" form:"playstore"`
	Appstote  string `json:"appstore" form:"appstore"`
}

type EditProductInput struct {
	Title     string `json:"title" form:"title"`
	Desc      string `json:"desc" form:"desc"`
	Image     string `json:"image" form:"image"`
	Playstore string `json:"playstore" form:"playstore"`
	Appstote  string `json:"appstore" form:"appstore"`
	ID        string `json:"prodID" form:"prodID"`
	UpdatedAt time.Time
}

type DeleteProductInput struct {
	ID string `json:"prodID" form:"prodID" binding:"required"`
}
