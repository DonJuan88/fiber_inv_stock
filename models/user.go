package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey"`
	FirstName string         `json:"firstname"`
	LastName  string         `json:"lastname"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
	IsAdmin   bool           `json:"admin" gorm:"default:false"`
	Active    bool           `json:"active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type UpdateUser struct {
	LastPassword         string `json:"last_password"`
	Password             string `json:"password" `
	PasswordConfirmation string `json:"password_confirmation" `
	Active               bool   `json:"active"`
}

type Register struct {
	ID                   uuid.UUID `json:"id" gorm:"primaryKey"`
	FirstName            string    `json:"firstname"`
	LastName             string    `json:"lastname"`
	Email                string    `json:"email" gorm:"unique"`
	Password             string    `json:"password"`
	PasswordConfirmation string    `json:"passwordconfirmation"`
	IsAdmin              bool      `json:"admin" gorm:"default:false"`
	Active               bool      `json:"active" gorm:"default:true"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
