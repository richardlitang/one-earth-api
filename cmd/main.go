package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heisenberglar/one-earth-api/database"
)

func main() {
	database.ConnectDatabase()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
