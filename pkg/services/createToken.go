package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	uuid2 "github.com/google/uuid"
	"io/ioutil"
	"math/rand"
	"musichain/pkg/dao"
	"musichain/pkg/domain"
	"musichain/pkg/http/requests"
	"musichain/pkg/http/response"
)

type TokenShare struct {
	Price            float64
	Div              float64
	InitialTktPool   uint
	RemainingTktPool uint
	MusicMediaId     uuid.UUID
}

type Token struct {
	NumShares        uint
	Price            float64
	Div              float64
	InitialTktPool   uint
	RemainingTktPool uint
	MusicMediaId     uuid.UUID
}

func CreateTokens(request requests.CreateTokenRequest, musicMedia *domain.MusicMedia) ([]TokenShare, error) {
	tokenList := make([]TokenShare, request.NumShares)

	for i := uint(0); i < request.NumShares; i++ {
		// TODO: Put the token in the Blockchain (not in the DB)
		token := TokenShare{
			Price:            request.Price,
			Div:              request.Div,
			InitialTktPool:   request.InitialTktPool,
			RemainingTktPool: request.InitialTktPool,
			MusicMediaId:     uuid.UUID(musicMedia.Id),
		}
		tokenList[i] = token
	}

	return tokenList, nil
}

func GetCreatedTokens(address string) ([]Token, error) {
	// This is a mock for temporary development
	tokenList := make([]Token, 3)

	inputString := "7f10426c-0936-43fd-b2a2-3f59f830fab9"
	id, _ := uuid.FromString(inputString)

	for i := uint(0); i < 3; i++ {
		// TODO: Put the token in the Blockchain (not in the DB)
		token := Token{
			NumShares:        uint(rand.Int()),
			Price:            float64(rand.Float64()),
			Div:              float64(rand.Float64()),
			InitialTktPool:   uint(rand.Int()),
			RemainingTktPool: uint(rand.Int()),
			MusicMediaId:     id,
		}
		tokenList[i] = token
	}

	return tokenList, nil
}

func GetMusicMedia(id string) (*response.MusicMediaResponse, error) {
	uuidMusicMedia, err := uuid.FromString(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	musicMedia, err := dao.GetMusicMedia(uuid2.UUID(uuidMusicMedia))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	mp3File, err := ioutil.ReadFile(musicMedia.Mp3Path)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Error reading MP3 file.")
	}

	imgFile, err := ioutil.ReadFile(musicMedia.ImgPath)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Error reading image file.")
	}

	musicMediaResponse := response.MusicMediaResponse{
		Id:             musicMedia.Id,
		Name:           musicMedia.Name,
		CreatorAddress: musicMedia.CreatorAddress,
		Mp3File:        mp3File,
		ImgFile:        imgFile,
	}

	return &musicMediaResponse, nil
}
