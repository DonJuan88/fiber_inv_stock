package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransferDetail struct {
	ID          uuid.UUID      `json:"id" gorm:"primaryKey"`
	TransferNo  string         `json:"transfer_no"`
	ProductCode string         `json:"product_code"`
	Qty         int64          `json:"qty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (td *TransferDetail) IncreaseDecreaseStock(tx *gorm.DB) (err error) {

	//bso : branch stock origin
	var bso BranchStock

	//decrease form origin branch
	if err := tx.Where("product_code = ? and branch_code = ?", bso.ProductCode, bso.BranchCode).First(&bso).Error; err != nil {
		return fmt.Errorf("source product not found: %w", err)
	}

	if bso.Stock < td.Qty {
		return fmt.Errorf("insufficient stock in source branch %s for product ID %s", bso.BranchCode, bso.ProductCode)
	}

	bso.Stock -= td.Qty
	if err := tx.Save(&bso).Error; err != nil {
		return fmt.Errorf("failed to decrease stock in source branch: %w", err)
	}

	// Increase stock in the destination branch

	//bsd : branch stock destiny
	var bsd BranchStock
	if err := tx.Where("branch_id = ? AND id = ?", bsd.BranchCode, bsd.ProductCode).First(&bsd).Error; err != nil {
		// If the product doesn't exist in the destination branch, create it
		newStock := BranchStock{
			BranchCode:  bsd.BranchCode,
			ProductCode: bsd.ProductCode,
			Stock:       td.Qty,
		}
		if err := tx.Create(&newStock).Error; err != nil {
			return fmt.Errorf("failed to create product in destination branch: %w", err)
		}
	} else {
		// If the product exists, just increase the stock
		bsd.Stock += td.Qty
		if err := tx.Save(&bsd).Error; err != nil {
			return fmt.Errorf("failed to increase stock in destination branch: %w", err)
		}
	}

	return nil
}
