package controllers

import (
	"backend/database"
	"backend/models"
	"backend/utility"

	"github.com/gofiber/fiber/v2"
)

func BusinessClients(c *fiber.Ctx) error {
	token, err := utility.VerifyAuth(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("client", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := database.GetUserFromToken(token)

	if user.AdminRole == 0 && user.BusinessRole == 0 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo(user.Name, "unauthorized")
		return c.JSON(fiber.Map{
			"message": "the user does not have a valid role for this request",
		})
	}

	var clients []models.Client

	if err := database.DB.Where("business_type = ?", 1).Find(&clients).Error; err != nil {
		utility.LogInfo("client", "noBusinessClients")
		return c.JSON(fiber.Map{
			"message": "Aucun clients",
		})
	}

	utility.LogInfo(user.Name, "businessClients")

	return c.JSON(clients)
}

func ResidentialClients(c *fiber.Ctx) error {
	token, err := utility.VerifyAuth(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("client", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := database.GetUserFromToken(token)

	if user.AdminRole == 0 && user.ResidentialRole == 0 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo(user.Name, "unauthorized")
		return c.JSON(fiber.Map{
			"message": "the user does not have a valid role for this request",
		})
	}

	var clients []models.Client

	if err := database.DB.Where("residential_type = ?", 1).Find(&clients).Error; err != nil {
		utility.LogInfo("client", "noResidentialClients")
		return c.JSON(fiber.Map{
			"message": "Aucun clients",
		})
	}

	utility.LogInfo(user.Name, "residentialClients")

	return c.JSON(clients)
}
