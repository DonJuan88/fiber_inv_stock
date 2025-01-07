package helper

import (
	"inv_fiber/config"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	combinedPassword := password + string(config.ENV.TOKEN_LOGIN)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(combinedPassword), bcrypt.DefaultCost)

	return string(passwordHash), err
}

func ComparePassword(hashedPassword, password string) bool {
 err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
 return err == nil
}

func HashEmail(email string) (string, error) {
	combinedEmail := email + string(config.ENV.TOKEN_EMAIL)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(combinedEmail), bcrypt.DefaultCost)

	return string(passwordHash), err
}