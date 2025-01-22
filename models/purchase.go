package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Purchase struct {
	ID            uuid.UUID        `json:"id" gorm:"primaryKey"`
	PurchaseNo    string           `json:"purchase_no"`
	PurchaseDate  time.Time        `json:"purchase_date"`
	BranchCode    string           `json:"foreignKey:BranchCode"`
	Supplier      string           `json:"foreignKey:SupplierCode"`
	ShippingCost  int64            `json:"shipping_price"`
	Tax1          int64            `json:"tax1"`
	Tax2          int64            `json:"tax2"`
	Total         int64            `json:"total"`
	UserID        string           `json:"foreignKey:UserID"`
	PaymentType   string           `json:"payment_type"`
	ShipStatus    string           `json:"shipping_status"`
	Reference     string           `json:"reference"`
	Notes         string           `json:"notes"`
	PaymentStatus bool             `jaon:"status gorm:default:false"`
	Details       []PurchaseDetail `gorm:"foreignKey:PurchaseNo"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	DeletedAt     gorm.DeletedAt   `json:"deleted_at" gorm:"index"`
}

func (p *Purchase) AfterCreate(tx *gorm.DB) (err error) {
	for _, detail := range p.Details {
		if err := detail.UpdateStock(tx); err != nil {
			return err
		}
	}
	return nil
}
func ValidatePurchase(purchase Purchase) error {
	if purchase.PurchaseNo == "" {
		return errors.New("purchase number is required")
	}

	layout := "2006-01-02T15:04:05"
	_, err := time.Parse(layout, purchase.PurchaseDate.String())
	if err != nil {
		return fmt.Errorf("invalid purchase date: %v", err)
	}

	return nil
}
