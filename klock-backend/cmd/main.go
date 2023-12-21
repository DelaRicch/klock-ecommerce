package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/DelaRicch/klock-ecommerce/klock-backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

)

func main() {
	database.ConnectDb()

	// Increase maximum payload limit
	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024, 
	})
	
	// set CORS
	app.Use(cors.New())
	// Set up all routes for the entire system
	setupRoutes(app)

	// Handle shutdown signals
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c

		// Call the CloseDb function before exiting
		database.CloseDb()

		// Stop the application gracefully
		os.Exit(0)
	}()

	// Start the Fiber app
	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}

}
