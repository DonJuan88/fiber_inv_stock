package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SaleDetail struct {
	ID          uuid.UUID      `json:"id" gorm:"primaryKey"`
	SaleNo      string         `json:"sale_no" gorm:"foreignKey:SaleNo"`
	SaleDate    time.Time      `json:"sale_date"`
	ProductCode string         `json:"product_code"`
	Qty         int64          `json:"qty"`
	Price       int64          `json:"price"`
	Discount    int64          `json:"discount"`
	SalePrice   int64          `json:"price_after_disc"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (sd *SaleDetail) DecreaseStock(tx *gorm.DB) (err error) {
	var bs BranchStock

	//check availably in branch
	if err := tx.Where("product_code = ? and branch_code = ?", bs.ProductCode, bs.BranchCode).First(&bs).Error; err != nil {

		return fmt.Errorf("Product %s not found ", err)
	}

	// Check if sufficient stock is available
	if bs.Stock < sd.Qty {
		return fmt.Errorf("insufficient stock for product ID %s", sd.ProductCode)
	}
	// Update existing stock
	bs.Stock -= sd.Qty
	if err := tx.Save(&bs).Error; err != nil {
		return fmt.Errorf("failed to update product stock: %w", err)
	}
	return nil
}
