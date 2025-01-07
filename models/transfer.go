package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transfer struct {

	ID			uuid.UUID     `json:"id" gorm:"primaryKey"`
	TransferNo    string `json:"transfer_no"`
	TransferDate time.Time `json:"transfer_date"`
	BranchOrigin  string `json:"branch_origin"`
	BranchDestiny string `json:"branch_destiny"`
	Reference     string `json:"reference"`
	Notes         string `json:"notes"`
	UserId        string `json:"user_id"`
	Cost          int64  `json:"cost"`
	CreatedAt time.Time       `json:"created_at"`
 	UpdatedAt time.Time       `json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
