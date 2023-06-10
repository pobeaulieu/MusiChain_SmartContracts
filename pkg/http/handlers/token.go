package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"musichain/pkg/dao"
	"musichain/pkg/http/requests"
	"musichain/pkg/services"
	"strconv"
)

func CreateTokens(c *fiber.Ctx) error {
	jwtToken, err := services.VerifyAuth(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "user must be authenticated")
	}

	claims := jwtToken.Claims.(*jwt.StandardClaims)
	uuidObj, err := uuid.Parse(claims.Issuer)
	creator, err := dao.GetCreatorFromId(uuidObj)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "user must exist and be a creator. ")
	}

	name := c.FormValue("name")
	numShares := c.FormValue("numShares")
	price := c.FormValue("price")
	div := c.FormValue("div")
	initialTktPool := c.FormValue("initialTktPool")
	mp3, errMp3 := c.FormFile("mp3")
	img, errImg := c.FormFile("img")

	if name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Name field is required.")
	}

	if numShares == "" {
		return fiber.NewError(fiber.StatusBadRequest, "numShares field is required.")
	}

	if price == "" {
		return fiber.NewError(fiber.StatusBadRequest, "price field is required.")
	}

	if div == "" {
		return fiber.NewError(fiber.StatusBadRequest, "div field is required.")
	}

	if initialTktPool == "" {
		return fiber.NewError(fiber.StatusBadRequest, "initialTktPool field is required.")
	}

	if errImg != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to retrieve img from request.")
	}

	if errMp3 != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to retrieve mp3 from request.")
	}

	numSharesFormatted, err := strconv.Atoi(numShares)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "NumShares must be an uint")
	}

	priceFormatted, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Price must be an float")
	}

	divFormatted, err := strconv.ParseFloat(div, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Price must be an float")
	}

	initialTktPoolFormatted, err := strconv.Atoi(initialTktPool)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "initialTktPool must be an uint")
	}

	request := requests.CreateTokenRequest{
		Creator:        creator,
		Name:           name,
		NumShares:      uint(numSharesFormatted),
		Price:          priceFormatted,
		Div:            divFormatted,
		InitialTktPool: uint(initialTktPoolFormatted),
		Mp3:            mp3,
		Img:            img,
	}

	musicMedia, errService := services.CreateMusicMedia(request)

	if errService != nil {
		return fiber.NewError(fiber.StatusBadRequest, errService.Error())
	}

	tokenList, errService := services.CreateTokens(request, musicMedia)

	if errService != nil {
		return fiber.NewError(fiber.StatusBadRequest, errService.Error())
	}

	return c.JSON(fiber.Map{
		"tokenList": tokenList,
	})

}
