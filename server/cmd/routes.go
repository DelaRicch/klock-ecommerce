package main

import (
	"github.com/DelaRicch/klock-ecommerce/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", handlers.Home)
	app.Get("/users", handlers.ListUsers)
	app.Post("/register", handlers.Register)
	app.Delete("/delete-users", handlers.DeleteAllUsers)
}