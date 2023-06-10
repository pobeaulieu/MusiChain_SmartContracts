package handlers

import (
	"github.com/gofiber/fiber/v2"
	"musichain/pkg/services"
)

func RegisterUser(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	username, ok := data["username"].(string)
	if !ok {
		return c.JSON(fiber.Map{
			"message": "error parsing username",
		})
	}

	password, ok := data["password"].(string)
	if !ok {
		return c.JSON(fiber.Map{
			"message": "error parsing password",
		})
	}

	isCreator, ok := data["isCreator"].(bool)
	if !ok {
		return c.JSON(fiber.Map{
			"message": "error parsing isCreator",
		})
	}

	user, err := services.RegisterUser(username, []byte(password), isCreator)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "error creating the user",
			"error":   err.Error(),
		})
	}

	return c.JSON(user)
}
