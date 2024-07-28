package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID       uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	StatusID    uint       `json:"status_id" binding:"required"`
	Status      TaskStatus `json:"status" gorm:"foreignKey:StatusID"`
}

type TaskResponse struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StatusID    int        `json:"status_id"`
	Status      TaskStatus `json:"status"`
}
