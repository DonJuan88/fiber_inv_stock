package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Branch struct {
	ID			uuid.UUID     `json:"id" gorm:"primaryKey"`
	BranchCode         string `json:"branch_code"`
	BranchName         string `json:"branch_name"`
	BranchAddress      string `json:"branch_address"`
	ContactPerson      string `json:"contact_person"`
	ContactPersonPhone string `json:"contact_person_phone"`
	Phone              string `json:"phone"`
	Active				bool  `json:"active"  gorm:"default:true"` 
	CreatedAt time.Time       `json:"created_at"`
 	UpdatedAt time.Time       `json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	
}
