package models

import "time"

type Log struct {
	Id       uint      `json:"id"`
	LogTime  time.Time `json:"logTime"`
	UserName string    `json:"userName"`
	Message  string    `json:"message"`
}
