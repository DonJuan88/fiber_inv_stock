package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID      `json:"id" gorm:"primaryKey"`
	ProductCode string         `json:"product_code"`
	Barcode1    string         `json:"barcode1"`
	Barcode2    string         `json:"barcode2"`
	ProductName string         `json:"name"`
	Description string         `json:"desc"`
	Category    string         `json:"category"`
	Brand       string         `json:"brand"`
	BasePrice   int64          `json:"baseprice"`
	SalePrice   int64          `json:"saleprice"`
	Unit        string         `json:"unit"`
	Active      bool           `json:"active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
