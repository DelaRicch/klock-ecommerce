package handlers

import (
	"fmt"

	"github.com/DelaRicch/klock-ecommerce/server/database"
	"github.com/DelaRicch/klock-ecommerce/server/lib"
	"github.com/DelaRicch/klock-ecommerce/server/models"
	"github.com/gofiber/fiber/v2"
)

func Home(ctx *fiber.Ctx) error {
	return ctx.SendString("This is Klock E-commerce web app")
}

func Register(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Check for the uniqueness of the email
	var existingUser models.User
	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	// Set the default role to "USER" if not specified
	if user.Role == "" {
		user.Role = "USER"
	}

	
    // Hash the password using Argon2
    hashedPassword, err := lib.HashPassword(user.Password)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Error hashing the password",
        })
    }
    user.Password = hashedPassword

	// Generate User ID
	user.UserID = lib.GenerateUserID()

// Create the User
	database.DB.Create(&user)
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("Successfylly registered %v", user.Name),
	})
}

func ListUsers(ctx *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Find(&users)
	return ctx.Status(fiber.StatusOK).JSON(users)
}

func DeleteAllUsers(ctx *fiber.Ctx) error {
	if err := database.DB.Exec("DELETE FROM users WHERE role = 'USER' ").Error; err != nil {
		return err
	}

	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "All users deleted",
	})
}
