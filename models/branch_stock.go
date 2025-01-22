package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BranchStock struct {
	ID          uuid.UUID      `json:"id" gorm:"primaryKey"`
	BranchCode  string         `json:"branch_code" gorm:"foreignKey:BranchCode"`
	ProductCode string         `json:"product_code" gorm:"foreignKey:ProductCode"`
	Stock       int64          `json:"stock"`
	MinStock    int64          `json:"min_stock"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
