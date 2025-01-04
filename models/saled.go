package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SaleDetail struct {
	
	ID			uuid.UUID     `json:"id" gorm:"primaryKey"`
	SaleNo   string    `json:"sale_no"`
	SaleDate time.Time `json:"sale_date"`
	ItemCode  string    `json:"code"`
	Qty       int64     `json:"qty"`
	Price     int64     `json:"price"`
	Discount  int64     `json:"discount"`
	SalePrice int64     `json:"price_after_disc"`
		CreatedAt time.Time       `json:"created_at"`
 	UpdatedAt time.Time       `json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
