package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sale struct {

	ID			uuid.UUID     `json:"id" gorm:"primaryKey"`
	SaleNo       string    `json:"sale_no"`
	SaleDate     time.Time `json:"sale_date"`
	Customer      string    `json:"customer"`
	ShippingCost  int64     `json:"shippingprice"`
	Tax1          int64     `json:"tax1"`
	Tax2          int64     `json:"tax2"`
	Total         int64     `json:"total"`
	AccountID     string    `json:"accid"`
	PaymentType   string    `json:"paymenttype"`
	Reference     string    `json:"reference"`
	Notes         string    `json:"notes"`
	PaymentStatus bool      `jaon:"status gorm:default:false"`
		CreatedAt time.Time       `json:"created_at"`
 	UpdatedAt time.Time       `json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	
}
