package services

import (
	"backend/dao"
	"backend/domain"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username string, password []byte, isCreator bool) (*domain.User, error) {
	hashPwd, _ := bcrypt.GenerateFromPassword(password, 14)

	return dao.InsertUser(username, hashPwd, isCreator)
}
