package models

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	Undone     Status = "undone"
	OnProgress Status = "on progress"
	Done       Status = "done"
)

type TodoItem struct {
	ID        uint      `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Status    Status    `json:"status"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type TodoItemQuery struct {
	Title   string    `json:"title"`
	Desc    string    `json:"desc"`
	Status  Status    `json:"status"`
	DueDate time.Time `json:"due_date"`
}

type UpdateTodoItemInput struct {
	Title   string    `json:"title"`
	Desc    string    `json:"desc"`
	Status  Status    `json:"status"`
	DueDate time.Time `json:"due_date"`
	ID      string    `json:"id"`
}

type UpdateTodoItemStatusInput struct {
	Status Status `json:"status"`
	ID     string `json:"id"`
}

type AddTodoItemInput struct {
	Title   string    `json:"title"`
	Desc    string    `json:"desc"`
	Status  Status    `json:"status"`
	DueDate time.Time `json:"due_date"`
}

type DeleteTodoItemInput struct {
	ID string `json:"id" form:"id" binding:"required"`
}
