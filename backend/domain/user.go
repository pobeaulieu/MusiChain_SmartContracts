package domain

import (
	"backend/database"
)

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password []byte `json:"-"`
}

func GetUser(username string) (*User, error) {
	var user User
	err := database.DB.Where("username = ?", username).First(&user).Error

	return &user, err
}
