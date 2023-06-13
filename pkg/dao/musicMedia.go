package dao

import (
	"fmt"
	"github.com/google/uuid"
	"musichain/pkg/database"
	"musichain/pkg/domain"
)

func InsertMusicMedia(id uuid.UUID, name string, creatorAddress string, mp3Path string, imgPath string) (*domain.MusicMedia, error) {
	musicMedia := domain.MusicMedia{
		Id:             id,
		Name:           name,
		CreatorAddress: creatorAddress,
		Mp3Path:        mp3Path,
		ImgPath:        imgPath,
	}

	tx := database.DB.Begin() // Start a transaction

	if err := tx.Create(&musicMedia).Error; err != nil {
		tx.Rollback() // Rollback the transaction if an error occurs
		return &domain.MusicMedia{}, err
	}

	tx.Commit() // Commit the transaction if all operations succeed

	return &musicMedia, nil
}

func GetMusicMedia(id uuid.UUID) (*domain.MusicMedia, error) {
	var musicMedia domain.MusicMedia
	fmt.Println("inside", id)
	database.DB.First(&musicMedia, id)

	return &musicMedia, nil
}
