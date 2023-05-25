package services

import (
	"backend/domain"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

func Login(c *fiber.Ctx, username string, password []byte) (*domain.User, error) {
	u, err := domain.GetUser(username)

	if err := bcrypt.CompareHashAndPassword(u.Password, password); err != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(u.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), //JWT valide 1 heure
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRETKEY")))

	if err != nil {
		return nil, err
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 1),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return u, nil

}

func VerifyAuth(c *fiber.Ctx) (*jwt.Token, error) {
	cookie := c.Cookies("jwt")

	return jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRETKEY")), nil
	})
}

func Logout(c *fiber.Ctx) {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
}
