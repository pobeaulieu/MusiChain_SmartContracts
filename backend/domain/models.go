package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `gorm:"type:uuid;primarykey"`
	Username string    `gorm:"unique"`
	Password []byte    `json:"-"`
}

type Creator struct {
	Id         uuid.UUID `gorm:"type:uuid;primarykey"`
	User       User      `gorm:"foreignkey:UserID"`
	MusicMedia []MusicMedia
	UserID     uuid.UUID
}

type MusicMedia struct {
	Id        uuid.UUID `gorm:"type:uuid;primarykey"`
	Name      string
	CreatorID uuid.UUID
	Mp3Path   string
	ImgPath   string
}
