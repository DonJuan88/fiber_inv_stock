package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransferDetail struct {

	ID			uuid.UUID     `json:"id" gorm:"primaryKey"`
	TransferNo string `json:"transfer_no"`
	ItemCode   string `json:"code"`
	Qty        int64  `json:"qty"`
	CreatedAt time.Time       `json:"created_at"`
 	UpdatedAt time.Time       `json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
