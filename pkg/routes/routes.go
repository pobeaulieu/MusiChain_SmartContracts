package routes

import (
	"github.com/gofiber/fiber/v2"
	"musichain/pkg/http/handlers"
)

func Setup(app *fiber.App) {
	app.Post("/api/registerUser", handlers.RegisterUser)
	app.Post("/api/login", handlers.Login)
	app.Post("/api/logout", handlers.Logout)
	app.Post("/api/createTokens", handlers.CreateTokens)
}
