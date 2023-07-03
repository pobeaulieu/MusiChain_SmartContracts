package services

import (
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"math/rand"
	"musichain/pkg/dao"
	"musichain/pkg/domain"
	"musichain/pkg/http/requests"
	"musichain/pkg/http/response"
	"strconv"
)

type TokenShare struct {
	Price            float64
	Div              float64
	InitialTktPool   uint
	RemainingTktPool uint
	MusicMediaId     uint
}

type Token struct {
	NumShares        uint
	Price            float64
	Div              float64
	InitialTktPool   uint
	RemainingTktPool uint
	MusicMediaId     uint
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
			MusicMediaId:     musicMedia.Id,
		}
		tokenList[i] = token
	}
	// TODO Deploy contract and mint on the adress where it is deployed

	// mintToken(int64(musicMedia.Id), int64(request.InitialTktPool), "adresseContrat", request.CreatorAddress)

	return tokenList, nil
}

func GetCreatedTokens(address string) ([]Token, error) {
	// This is a mock for temporary development
	tokenList := make([]Token, 3)
	// TODO: Nico ajoute lappel au blockchain pour obtenir tous les tokens crees avec address en parametre
	for i := uint(0); i < 3; i++ {

		token := Token{
			NumShares:        uint(rand.Int()),
			Price:            float64(rand.Float64()),
			Div:              float64(rand.Float64()),
			InitialTktPool:   uint(rand.Int()),
			RemainingTktPool: uint(rand.Int()),
			MusicMediaId:     1,
		}
		tokenList[i] = token
	}

	return tokenList, nil
}

func GetMusicMedia(id string) (*response.MusicMediaResponse, error) {

	musicMediaId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	musicMedia, err := dao.GetMusicMedia(uint(musicMediaId))
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
