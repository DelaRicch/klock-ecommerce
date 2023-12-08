package main

import (
	"github.com/DelaRicch/klock-ecommerce/server/handlers"
	// jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Get("/users", handlers.ListUsers)
	app.Get("/refresh-token", handlers.RequestNewToken)
	app.Get("/user-profile", handlers.GetUserProfile)
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Post("/social-login", handlers.SocialLogin)
	app.Delete("/delete-users", handlers.DeleteAllUsers)

	// Products routes 
	app.Post("/add-product", handlers.AddNewProduct)
	app.Get("/products", handlers.ListProducts)
	app.Delete("/delete-products", handlers.DeleteAllProducts)

	// Protected Routes
	// ProtectedRoutes := app.Group("/")
	// ProtectedRoutes.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	// }))

	// ProtectedRoutes.Get("/", handlers.Home)

	app.Get("/", handlers.Home)
}
