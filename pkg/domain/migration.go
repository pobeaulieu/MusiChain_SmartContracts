package domain

import (
	"musichain/pkg/database"
)

func AutoMigrate() error {
	return database.DB.AutoMigrate(&User{}, &Creator{}, &MusicMedia{})
}
