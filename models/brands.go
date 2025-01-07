package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Brands struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey"`
	BrandCode string         `json:"brand_code"`
	BrandName string         `json:"brand_name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
