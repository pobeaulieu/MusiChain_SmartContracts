package domain

type MusicMedia struct {
	Id             uint `gorm:"primarykey"`
	Name           string
	CreatorAddress string
	Mp3Path        string
	ImgPath        string
}
