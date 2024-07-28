package model

import (
	"gorm.io/gorm"
)

type TaskStatus struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	StatusName string `json:"status_name"`
}

type TaskStatusRequest struct {
	StatusName string `json:"status_name"`
}
type TaskStatusResponse struct {
	ID         int    `json:"id"`
	StatusName string `json:"status_name"`
}
