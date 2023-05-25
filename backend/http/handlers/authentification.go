package handlers

import (
	"backend/services"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := services.Login(c, data["username"], []byte(data["password"]))

	if err != nil {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	services.Logout(c)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
