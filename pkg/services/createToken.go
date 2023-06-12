package services

import (
	"fmt"
	"github.com/google/uuid"
	"musichain/pkg/domain"
	"musichain/pkg/http/requests"
)

type Token struct {
	Price            float64
	Div              float64
	InitialTktPool   uint
	RemainingTktPool uint
	MusicMediaId     uuid.UUID
}

func CreateTokens(request requests.CreateTokenRequest, musicMedia *domain.MusicMedia) ([]Token, error) {
	tokenList := make([]Token, request.NumShares)

	for i := uint(0); i < request.NumShares; i++ {
		// TODO: Put the token in the Blockchain (not in the DB)
		token := Token{
			Price:            request.Price,
			Div:              request.Div,
			InitialTktPool:   request.InitialTktPool,
			RemainingTktPool: request.InitialTktPool,
			MusicMediaId:     musicMedia.Id,
		}
		tokenList[i] = token
	}

	fmt.Println(tokenList)

	return tokenList, nil
}

func GetCreatedTokens(address string) ([]Token, error) {
	// This is a mock for temporary development
	tokenList := make([]Token, 3)

	inputString := "082c6a7f-3dce-4523-9f45-64df69da4e41"
	uuid := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(inputString))

	for i := uint(0); i < 3; i++ {
		// TODO: Put the token in the Blockchain (not in the DB)
		token := Token{
			Price:            2,
			Div:              0.1,
			InitialTktPool:   100000,
			RemainingTktPool: 54590,
			MusicMediaId:     uuid,
		}
		tokenList[i] = token
	}

	fmt.Println(tokenList)

	return tokenList, nil
}
