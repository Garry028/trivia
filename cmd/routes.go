package main

import (
	"github.com/Garry028/trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)
}
