package main

import (
	"github.com/DelaRicch/klock-ecommerce/server/database"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	// set CORS
	app.Use(cors.New())

	// Set up all routes for the entire system
	setupRoutes(app)



	app.Listen(":3000")
}
