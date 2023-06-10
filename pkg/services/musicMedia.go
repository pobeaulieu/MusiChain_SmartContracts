package services

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"musichain/pkg/dao"
	"musichain/pkg/domain"
	"musichain/pkg/http/requests"
	"os"
)

func CreateMusicMedia(request requests.CreateTokenRequest) (*domain.MusicMedia, error) {
	// Generate a unique ID for the MusicMedia
	musicMediaID := uuid.New()

	// Save the mp3 file to the specified path with the generated ID as the file name
	mp3Path := fmt.Sprintf("./database/mp3/%s.mp3", musicMediaID)
	err := saveFile(request.Mp3, mp3Path)
	if err != nil {
		return nil, err
	}

	// Save the mp3 file to the specified path with the generated ID as the file name
	imgPath := fmt.Sprintf("./database/img/%s.png", musicMediaID)
	err = saveFile(request.Img, imgPath)
	if err != nil {
		return nil, err
	}

	// Insert the MusicMedia into the database
	return dao.InsertMusicMedia(musicMediaID, request.Name, request.Creator, mp3Path, imgPath)
}

func saveFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}
