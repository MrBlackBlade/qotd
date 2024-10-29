package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/MrBlackBlade/qotd/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.QuoteOfTheDay)

	app.Get("/all", handlers.ListQuotes)

	app.Post("/quotes", handlers.CreateQuotes)

	/*app.Post("/multiquotes", handlers.CreateQuotes)*/
}
