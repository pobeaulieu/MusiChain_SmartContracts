package domain

import (
	"backend/database"
)

func AutoMigrate() error {
	return database.DB.AutoMigrate(&User{}, &Creator{}, &MusicMedia{})
}
