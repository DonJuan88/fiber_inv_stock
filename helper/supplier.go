package helper

import (
	"inv_fiber/models"

	"gorm.io/gorm"
)

// check double code
func CheckSupplierExists(db *gorm.DB, code string) (bool, error) {
	var supplier models.Supplier
	result := db.Where("code = ?", code).First(&supplier)

	// If record found, return true
	if result.RowsAffected > 0 {
		return true, nil
	}

	// If no record found, return false
	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	}

	// For other database errors, return false and error
	return false, result.Error
}
