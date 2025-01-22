package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	ID           uuid.UUID      `json:"id" gorm:"primaryKey"`
	NotifId      string         `json:"notif_id"`
	UserID       string         `json:"user_id" gorm:"foreignKey:UserID"`
	NotifMessage string         `json:"notif_message"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
