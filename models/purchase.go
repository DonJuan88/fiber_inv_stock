package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Purchase struct {
	ID            uuid.UUID      `json:"id" gorm:"primaryKey"`
	PurchaseNo    string         `json:"purchase_no"`
	PurchaseDate  time.Time      `json:"purchase_date"`
	BranchCode    string         `json:"branch"`
	Supplier      string         `json:"supplier"`
	ShippingCost  int64          `json:"shippingprice"`
	Tax1          int64          `json:"tax1"`
	Tax2          int64          `json:"tax2"`
	Total         int64          `json:"total"`
	AccountID     string         `json:"accid"`
	PaymentType   string         `json:"paymenttype"`
	ShipStatus    string         `json:"shippingstatus"`
	Reference     string         `json:"reference"`
	Notes         string         `json:"notes"`
	PaymentStatus bool           `jaon:"status gorm:default:false"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
