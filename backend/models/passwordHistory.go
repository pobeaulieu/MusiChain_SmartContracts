package models

import (
	"time"
)

type PasswordHistory struct {
	Id          uint      `gorm:"primaryKey"`
	UserEmail   string    `json:"userEmail"`
	CreatedTime time.Time `json:"createdTime"`
	Password    []byte    `json:"-"`
}
