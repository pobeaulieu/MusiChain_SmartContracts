package domain

import (
	"github.com/google/uuid"
)

type MusicMedia struct {
	Id             uuid.UUID `gorm:"type:uuid;primarykey"`
	Name           string
	CreatorAddress string
	Mp3Path        string
	ImgPath        string
}
