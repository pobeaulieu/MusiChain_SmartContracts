package handlers

import (
	"backend/services"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := services.RegisterUser(data["username"], []byte(data["password"]))

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "error creating the user",
			"error":   err.Error(),
		})
	}

	return c.JSON(user)
}
