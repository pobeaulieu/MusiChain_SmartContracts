package controllers

import (
	"backend/database"
	"backend/models"
	"backend/utility"

	"github.com/gofiber/fiber/v2"
)

func User(c *fiber.Ctx) error {
	token, err := utility.VerifyAuth(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("user", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := database.GetUserFromToken(token)

	utility.LogInfo(user.Name, "user")

	return c.JSON(user)
}

func Users(c *fiber.Ctx) error {
	token, err := utility.VerifyAuth(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("user", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := database.GetUserFromToken(token)

	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		utility.LogInfo("user", "noUsers")
		return c.JSON(fiber.Map{
			"message": "Aucun utilisateurs",
		})
	}

	if user.AdminRole == 0 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo(user.Name, "unauthorized")
		return c.JSON(fiber.Map{
			"message": "the user does not have a valid admin role",
		})
	}

	utility.LogInfo(user.Name, "users")

	return c.JSON(users)
}
