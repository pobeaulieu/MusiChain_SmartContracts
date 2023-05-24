package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/users", controllers.Users)
	app.Get("/api/businessclients", controllers.BusinessClients)
	app.Get("/api/residentialclients", controllers.ResidentialClients)
	app.Post("/api/unblock", controllers.Unblock)
	app.Post("/api/maxloginattempts", controllers.MaxLoginAttempts)
	app.Post("/api/logintimeinterval", controllers.LoginTimeInterval)
	app.Post("/api/passwordpolicy", controllers.PasswordPolicy)
	app.Get("/api/getpasswordpolicy", controllers.GetPasswordPolicy)
	app.Get("/api/getloginpolicy", controllers.GetLoginPolicy)
	app.Post("/api/userrole", controllers.UserRole)
	app.Post("/api/modifypassword", controllers.ModifyPassword)
	app.Post("/api/block", controllers.Block)
	app.Post("/api/login2fa", controllers.Login2FA)
	app.Post("/api/activate2fa", controllers.Activate2FA)
	app.Get("/api/getcode", controllers.GetCode)
}
