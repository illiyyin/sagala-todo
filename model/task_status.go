package model

import (
	"gorm.io/gorm"
)

type TaskStatus struct {
	gorm.Model
	StatusName string `json:"status_name"`
	StatusID   int    `json:"id" gorm:"unique;primaryKey"`
}
