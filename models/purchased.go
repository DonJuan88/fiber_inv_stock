package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PurchaseDetail struct {
	ID            uuid.UUID      `json:"id" gorm:"primaryKey"`
	PurchaseNo    string         `json:"purchase_no"`
	PurchaseDate  time.Time      `json:"purchase_date"`
	ItemCode      string         `json:"code"`
	Qty           int64          `json:"qty"`
	BasePrice     int64          `json:"baseprice"`
	Discount      int64          `json:"discount"`
	PurchasePrice int64          `json:"purchaseprice"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
