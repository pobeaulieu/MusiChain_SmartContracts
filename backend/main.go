package main

import (
	"backend/database"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

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
