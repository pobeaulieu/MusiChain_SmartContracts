package routes

import (
	"github.com/gofiber/fiber/v2"
	"musichain/pkg/http/handlers"
)

func Setup(app *fiber.App) {
	app.Post("/api/createTokens", handlers.CreateTokens)
}
