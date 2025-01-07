package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Categories struct {
	ID           uuid.UUID      `json:"id" gorm:"primaryKey"`
	CategoryCode string         `json:"category_code"`
	CategoryName string         `json:"category_name"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
