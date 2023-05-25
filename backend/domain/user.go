package domain

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password []byte `json:"-"`
}

func (u *User) Login(c *fiber.Ctx, password []byte) error {
	if err := bcrypt.CompareHashAndPassword(u.Password, password); err != nil {
		return fmt.Errorf("incorrect password")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(u.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), //JWT valide 1 heure
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRETKEY")))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 1),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return nil

}

func VerifyAuth(c *fiber.Ctx) (*jwt.Token, error) {
	cookie := c.Cookies("jwt")

	return jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRETKEY")), nil
	})
}

func GetUser(username string) (*User, error) {
	var user User
	err := DB.Where("username = ?", username).First(&user).Error

	return &user, err
}

func CreateUser(username string, password []byte) (*User, error) {
	hashPwd, _ := bcrypt.GenerateFromPassword(password, 14)

	user := User{
		Username: username,
		Password: hashPwd,
	}

	err := DB.Create(&user).Error

	return &user, err
}
