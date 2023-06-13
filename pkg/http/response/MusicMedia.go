package response

import "github.com/google/uuid"

type MusicMediaResponse struct {
	Id             uuid.UUID `json:"Id"`
	Name           string    `json:"Name"`
	CreatorAddress string    `json:"CreatorAddress"`
	Mp3File        []byte    `json:"Mp3File"`
	ImgFile        []byte    `json:"ImgFile"`
}
