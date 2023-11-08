package main

import (
	"github.com/DelaRicch/klock-ecommerce/server/handlers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	// app.Get("/", handlers.Home)
	app.Get("/users", handlers.ListUsers)
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Delete("/delete-users", handlers.DeleteAllUsers)

	// Protected Routes
	ProtectedRoutes := app.Group("/")
	ProtectedRoutes.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	ProtectedRoutes.Get("/", handlers.Home)
}