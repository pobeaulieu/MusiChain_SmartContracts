package services

import (
	"backend/database"
	"backend/domain"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username string, password []byte) (*domain.User, error) {
	hashPwd, _ := bcrypt.GenerateFromPassword(password, 14)

	user := domain.User{
		Username: username,
		Password: hashPwd,
	}

	err := database.DB.Create(&user).Error

	return &user, err
}
