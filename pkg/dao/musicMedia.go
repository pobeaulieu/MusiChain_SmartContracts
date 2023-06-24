package dao

import (
	"fmt"
	"musichain/pkg/database"
	"musichain/pkg/domain"
	"strconv"
)

func InsertMusicMedia(name string, creatorAddress string) (*domain.MusicMedia, error) {
	tx := database.DB.Begin() // Start a transaction
	musicMediaId := newMusicMediaId()
	// Save the mp3 file to the specified path with the generated ID as the file name
	mp3Path := fmt.Sprintf("./database/mp3/%s.mp3", strconv.Itoa(int(musicMediaId)))
	// Save the mp3 file to the specified path with the generated ID as the file name
	imgPath := fmt.Sprintf("./database/img/%s.png", strconv.Itoa(int(musicMediaId)))
	musicMedia := domain.MusicMedia{
		Id:             musicMediaId,
		Name:           name,
		CreatorAddress: creatorAddress,
		Mp3Path:        mp3Path,
		ImgPath:        imgPath,
	}

	if err := tx.Create(&musicMedia).Error; err != nil {
		tx.Rollback() // Rollback the transaction if an error occurs
		return &domain.MusicMedia{}, err
	}

	tx.Commit() // Commit the transaction if all operations succeed

	return &musicMedia, nil
}

func newMusicMediaId() uint {
	var maxID uint
	database.DB.Model(&domain.MusicMedia{}).Select("MAX(id)").Scan(&maxID)
	return maxID + 1

}

func GetMusicMedia(id uint) (*domain.MusicMedia, error) {
	var musicMedia domain.MusicMedia
	database.DB.First(&musicMedia, id)

	return &musicMedia, nil
}
