package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heisenberglar/one-earth-api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/quests", handlers.GetQuests)
	app.Get("/quests/:id", handlers.GetQuest)
	app.Post("/quests", handlers.CreateQuest)
	app.Put("/quests/:id", handlers.UpdateQuest)
	app.Delete("/quests/:id", handlers.DeleteQuest)
}
