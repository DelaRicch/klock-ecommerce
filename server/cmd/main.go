package main

import (
	"os"
	"syscall"
	"os/signal"

	"github.com/DelaRicch/klock-ecommerce/server/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

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
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
