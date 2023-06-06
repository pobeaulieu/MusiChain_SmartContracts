package dao

import (
	"backend/database"
	"backend/domain"
	"github.com/google/uuid"
)

func GetUserFromUsername(username string) (*domain.User, error) {
	var user domain.User
	err := database.DB.Where("username = ?", username).First(&user).Error

	return &user, err
}

func GetUserFromId(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	err := database.DB.Where("id = ?", id).First(&user).Error

	return &user, err
}

func GetCreatorFromId(id uuid.UUID) (*domain.Creator, error) {
	var creator domain.Creator
	err := database.DB.Where("id = ?", id).First(&creator).Error

	return &creator, err
}

func InsertUser(username string, hashPwd []byte, isCreator bool) (*domain.User, error) {

	id := uuid.New()

	user := domain.User{
		Id:       id,
		Username: username,
		Password: hashPwd,
	}

	tx := database.DB.Begin() // Start a transaction

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback() // Rollback the transaction if an error occurs
		return nil, err
	}

	if isCreator {
		creator := domain.Creator{
			Id:         id,
			User:       user,
			MusicMedia: []domain.MusicMedia{}, // Initialize an empty slice for associated music media
		}
		if err := tx.Create(&creator).Error; err != nil {
			tx.Rollback() // Rollback the transaction if an error occurs
			return nil, err
		}
	}

	tx.Commit() // Commit the transaction if all operations succeed

	return &user, nil
}
