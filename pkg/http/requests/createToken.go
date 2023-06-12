package requests

import (
	"mime/multipart"
)

type CreateTokenRequest struct {
	CreatorAddress string
	Name           string
	NumShares      uint
	Price          float64
	Div            float64
	InitialTktPool uint
	Mp3            *multipart.FileHeader
	Img            *multipart.FileHeader
}
