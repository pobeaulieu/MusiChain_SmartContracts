package services

import (
	"golang.org/x/crypto/bcrypt"
	"musichain/pkg/dao"
	"musichain/pkg/domain"
)

func RegisterUser(username string, password []byte, isCreator bool) (*domain.User, error) {
	hashPwd, _ := bcrypt.GenerateFromPassword(password, 14)

	return dao.InsertUser(username, hashPwd, isCreator)
}
