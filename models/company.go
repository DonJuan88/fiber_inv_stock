package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID                 uuid.UUID      `json:"id" gorm:"primaryKey"`
	Name               string         `json:"company_name"`
	ContactPerson      string         `json:"cp"`
	ContactPersonPhone string         `json:"cp_phone"`
	Email              string         `json:"email"`
	Phone              string         `json:"phone"`
	Address            string         `json:"address"`
	City               string         `json:"city"`
	State              string         `json:"state"`
	PostalCode         string         `json:"postalcode"`
	Country            string         `json:"country"`
	LocationID         string         `json:"location"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
