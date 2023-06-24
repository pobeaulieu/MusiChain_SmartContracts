package services

import (
	"io"
	"mime/multipart"
	"musichain/pkg/dao"
	"musichain/pkg/domain"
	"musichain/pkg/http/requests"
	"os"
)

func CreateMusicMedia(request requests.CreateTokenRequest) (musicMedia *domain.MusicMedia, err error) {
	if musicMedia, err = dao.InsertMusicMedia(request.Name, request.CreatorAddress); err != nil {
		return
	}

	if err = saveFile(request.Mp3, musicMedia.Mp3Path); err != nil {
		return
	}

	if err = saveFile(request.Img, musicMedia.ImgPath); err != nil {
		return nil, err
	}

	return
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
