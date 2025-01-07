package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID            uuid.UUID      `json:"id" gorm:"primaryKey"`
	Name          string         `json:"name"`
	EmailPhone    string         `json:"email_phone"`
	MyMessage     string         `json:"message"`
	ReadingStatus bool           `json:"status"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
