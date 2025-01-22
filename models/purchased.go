package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PurchaseDetail struct {
	ID            uuid.UUID      `json:"id" gorm:"primaryKey"`
	PurchaseNo    string         `json:"purchase_no"`
	PurchaseDate  time.Time      `json:"purchase_date"`
	ProductCode   string         `json:"product_code" gorm:"foreignKey:ProductCode"`
	Qty           int64          `json:"qty"`
	Price         int64          `json:"price"`
	Discount      int64          `json:"discount"`
	PurchasePrice int64          `json:"purchase_price"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (pd *PurchaseDetail) UpdateStock(tx *gorm.DB) (err error) {
	var bs BranchStock
	if err := tx.Where("product_code = ? and branch_code = ?", bs.ProductCode, bs.BranchCode).First(&bs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create new stock record if not found
			newStock := BranchStock{
				BranchCode:  bs.BranchCode,
				ProductCode: pd.ProductCode,
				Stock:       pd.Qty,
			}
			return tx.Create(&newStock).Error
		}
		return err
	}

	// Update existing stock
	bs.Stock += pd.Qty
	return tx.Save(&bs).Error
}
