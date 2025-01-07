package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductImage struct {
	ID          uuid.UUID      `json:"id" gorm:"primaryKey"`
	ProductCode string         `json:"product_code"`
	FileName    string         `json:"product_image"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type ImageUploads struct {
	ProductCode string `json:"code"`
	FileName    string `json:"image"`
}
