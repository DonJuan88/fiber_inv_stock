package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID            uuid.UUID      `json:"id" gorm:"primaryKey"`
	Code          string         `json:"code"`
	Name          string         `json:"name"`
	ContactPerson string         `json:"cp"`
	Email         string         `json:"email"`
	Phone         string         `json:"phone"`
	Address       string         `json:"address"`
	City          string         `json:"city"`
	State         string         `json:"state"`
	PostalCode    string         `json:"postalcode"`
	Country       string         `json:"country"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
