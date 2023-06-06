package main

import (
	"backend/database"
	"backend/domain"
	"backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	if err := domain.AutoMigrate(); err != nil {
		panic("could not migrate the database")
	}

	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024, // 50 MB in bytes
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Load SSL certificate and key
	certFile := "./.cert/cert2.pem"
	keyFile := "./.cert/key2.pem"

	routes.Setup(app)

	// Start HTTPS server
	err := app.ListenTLS(":8000", certFile, keyFile)

	if err != nil {
		panic(err)
	}

}
