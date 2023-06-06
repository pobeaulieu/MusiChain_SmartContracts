package routes

import (
	"backend/http/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/registerUser", handlers.RegisterUser)
	app.Post("/api/login", handlers.Login)
	app.Post("/api/logout", handlers.Logout)
	app.Post("/api/createTokens", handlers.CreateTokens)
}
