package handlers

import (
	"fmt"
	"time"

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
			"success": false,
		})
	}

	// Check for the uniqueness of the email
	var existingUser models.User
	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Email already exists",
			"success": false,
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
			"success": false,
		})
	}
	user.Password = hashedPassword

	// Generate User ID
	user.UserID = lib.GenerateUserID()

	// Create the User
	database.DB.Create(&user)

	// Generate JWT
	token, exp, err := lib.CreateJwtToken(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate JWT",
			"success": false,
		})
	}

	// Set token to the users' cookies for future requests
	ctx.Cookie(&fiber.Cookie{
		Name:        "access_token",
		Value:       fmt.Sprintf("Bearer %v", token),
		Expires:     time.Now().Add(time.Minute * 30),
		Secure:      true,
		HTTPOnly:    true,
		SessionOnly: true,
	})

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success":      true,
		"exp":          exp,
		"message":      fmt.Sprintf("Successfylly registered %v", user.Name),
		"access_token": token,
	})
}

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
	})

	if err := ctx.BodyParser(loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// Retrieve the user with the given email
	var user models.User
	result := database.DB.Where("email = ?", loginRequest.Email).First(&user)
	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// Verify the password using Argon2id
	if !lib.VerifyPassword(user.Password, loginRequest.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// Generate JWt
	token, exp, err := lib.CreateJwtToken(&user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate JWT",
			"success": false,
		})
	}

	// Set token to the users' cookies for future requests
	ctx.Cookie(&fiber.Cookie{
		Name:        "access_token",
		Value:       fmt.Sprintf("Bearer %v", token),
		Expires:     time.Now().Add(time.Minute * 30),
		Secure:      true,
		HTTPOnly:    true,
		SessionOnly: true,
	})

	return ctx.JSON(fiber.Map{
		"success":      true,
		"exp":          exp,
		"message":      fmt.Sprintf("Welcome %v", user.Name),
		"access_token": token,
	})

}

func ListUsers(ctx *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Find(&users)
	return ctx.Status(fiber.StatusOK).JSON(users)
}

func DeleteAllUsers(ctx *fiber.Ctx) error {
	if err := database.DB.Exec("DELETE FROM users WHERE role = 'ADMIN' ").Error; err != nil {
		return err
	}

	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "All users deleted",
	})
}
