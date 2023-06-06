package dao

import (
	"backend/database"
	"backend/domain"
	"github.com/google/uuid"
)

func InsertMusicMedia(id uuid.UUID, name string, creator *domain.Creator, mp3Path string, imgPath string) (*domain.MusicMedia, error) {
	musicMedia := domain.MusicMedia{
		Id:        id,
		Name:      name,
		CreatorID: creator.Id,
		Mp3Path:   mp3Path,
		ImgPath:   imgPath,
	}

	tx := database.DB.Begin() // Start a transaction

	if err := tx.Create(&musicMedia).Error; err != nil {
		tx.Rollback() // Rollback the transaction if an error occurs
		return &domain.MusicMedia{}, err
	}

	tx.Commit() // Commit the transaction if all operations succeed

	return &musicMedia, nil
}
