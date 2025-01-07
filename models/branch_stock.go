package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BranchStockPrice struct {
	ID			uuid.UUID     `json:"id" gorm:"primaryKey"`
	BranchCode  string `json:"branch_code"`
	ProductCode string `json:"code"`
	Barcode1    string `json:"barcode1"`
	Barcode2    string `json:"barcode2"`
	BasePrice       int64  `json:"baseprice"`
	SalePrice int64 `json:"saleprice"`
	Stock       int64  `json:"stock"`
	MinStock int64 `json:"min_stock"`
	CreatedAt time.Time       `json:"created_at"`
 	UpdatedAt time.Time       `json:"updated_at"`
 	DeletedAt gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
}
