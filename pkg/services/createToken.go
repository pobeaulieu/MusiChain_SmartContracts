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
