package requests

import (
	"backend/domain"
	"mime/multipart"
)

type CreateTokenRequest struct {
	Creator        *domain.Creator
	Name           string
	NumShares      uint
	Price          float64
	Div            float64
	InitialTktPool uint
	Mp3            *multipart.FileHeader
	Img            *multipart.FileHeader
}
