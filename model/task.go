package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          string     `json:"id" gorm:"type:uuid;primaryKey"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	StatusID    int        `json:"status_id" binding:"required"`
	Status      TaskStatus `json:"status" gorm:"foreignKey:StatusID"`
}
