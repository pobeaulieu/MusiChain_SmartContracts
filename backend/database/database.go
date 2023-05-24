package database

import (
	"backend/models"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load("environment.env")
	dbURL := os.Getenv("DB_URL")
	connection, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Client{}, &models.PasswordPolicy{}, &models.LoginPolicy{}, &models.Log{}, &models.PasswordHistory{})
}

func GetUserFromToken(token *jwt.Token) *models.User {
	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	DB.Where("id = ?", claims.Issuer).First(&user)

	return &user
}

func GetPasswordPolicy(user models.User) *models.PasswordPolicy {
	var policy models.PasswordPolicy
	DB.Where("id = ?", user.PasswordPolicyId).First(&policy)

	return &policy
}

func GetLoginPolicy(user *models.User) (*models.LoginPolicy, error) {
	var policy models.LoginPolicy
	err := DB.Where("id = ?", user.LoginPolicyId).First(&policy).Error

	return &policy, err
}

func GetUserFromEmail(email string) (*models.User, error) {
	var user models.User
	err := DB.Where("email = ?", email).First(&user).Error

	return &user, err
}
