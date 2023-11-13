package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/DelaRicch/klock-ecommerce/server/database"
	"github.com/DelaRicch/klock-ecommerce/server/lib"
	"github.com/DelaRicch/klock-ecommerce/server/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const errorRfTokenMsg string = "Error generating refresh token"

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
	result = database.DB.Create(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Internal server error",
			"success": false,
		})
	}

	// Generate JWT
	refreshTkn, token, exp, err := lib.CreateJwtToken(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errorRfTokenMsg,
			"success": false,
		})
	}

	// Set token to the users' cookies for future requests
	ctx.Cookie(&fiber.Cookie{
		Name:        "access_token",
		Value:       token,
		Expires:     time.Now().Add(time.Hour * 1),
		Secure:      true,
		HTTPOnly:    true,
		SessionOnly: true,
	})

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success":       true,
		"exp":           exp,
		"message":       fmt.Sprintf("Successfully registered %v", user.Name),
		"access_token":  token,
		"refresh_token": refreshTkn,
	})
}

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(models.User)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// extend access token duration if remember me is true
	var tokenExpiry time.Time
	if loginRequest.RememberMe == "true" {
		tokenExpiry = time.Now().Add(time.Hour * 24 * 30)
	} else {
		tokenExpiry = time.Now().Add(time.Hour * 1)
	}

	// Retrieve the user with the given email
	var user models.User
	fmt.Println(loginRequest, "Login Request")
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
	refreshTkn, token, _, err := lib.CreateJwtToken(&user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errorRfTokenMsg,
			"success": false,
		})
	}

	// Set token to the users' cookies for future requests
	ctx.Cookie(&fiber.Cookie{
		Name:        "access_token",
		Value:       token,
		Expires:     tokenExpiry,
		Secure:      true,
		HTTPOnly:    true,
		SessionOnly: true,
	})

	return ctx.JSON(fiber.Map{
		"success":       true,
		"exp":           tokenExpiry.Unix(),
		"message":       fmt.Sprintf("Welcome %v", user.Name),
		"access_token":  token,
		"refresh_token": refreshTkn,
	})

}

// Social login
func SocialLogin(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}

	fmt.Println(user, "Social Login Request")

	// Check for the uniqueness of the email
	email := user.Email
	var existingUser models.User
	result := database.DB.Where("email = ?", email).First(&existingUser)
	if result.RowsAffected > 0 {	
		if existingUser.Password != "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Social login not allowed for this account",
				"success": false,
			})
		}

		if existingUser.SocialId != user.SocialId {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized social login",
				"success": false,
			})
		}

		// Generate JWT
		refreshTkn, token, exp, err := lib.CreateJwtToken(user)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": errorRfTokenMsg,
				"success": false,
			})
		}

		// Set token to the users' cookies for future requests
		ctx.Cookie(&fiber.Cookie{
			Name:        "access_token",
			Value:       token,
			Expires:     time.Now().Add(time.Hour * 1),
			Secure:      true,
			SessionOnly: true,
		})

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success":       true,
			"exp":           exp,
			"message":       fmt.Sprintf("Welcome %v", user.Name),
			"access_token":  token,
			"refresh_token": refreshTkn,
		})

	}

	// Create a new user if the user doesn't exist
	user.UserID = lib.GenerateUserID()

	// Set the default role to "USER" if not specified
	if user.Role == "" {
		user.Role = "USER"
	}

	// Add the user to the DB
	result = database.DB.Create(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Internal server error",
			"success": false,
		})
	}

	// Generate JWT
	refreshTkn, token, exp, err := lib.CreateJwtToken(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errorRfTokenMsg,
			"success": false,
		})
	}

	// Set token to the users' cookies for future requests
	ctx.Cookie(&fiber.Cookie{
		Name:        "access_token",
		Value:       token,
		Expires:     time.Now().Add(time.Hour * 1),
		Secure:      true,
		SessionOnly: true,
	})

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success":       true,
		"exp":           exp,
		"message":       fmt.Sprintf("Successfully registered %v", user.Name),
		"access_token":  token,
		"refresh_token": refreshTkn,
	})
}

// Request new token using refresh token
func RequestNewToken(ctx *fiber.Ctx) error {
	refreshToken := ctx.Get("Authorization")

	if refreshToken == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Refresh token required",
		})
	}

	// Extract the token part from "Bearer <token>"
	tokenParts := strings.Split(refreshToken, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid refresh token format",
		})
	}
	refreshToken = tokenParts[1]

	// Parse and validate the refresh token
	tokenByte, err := jwt.Parse(refreshToken, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		// Use the same secret key for verifying the refresh token
		return []byte("rf_secret"), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid refresh token",
		})
	}

	// Extract claims from the refresh token
	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Invalid token claim"})
	}

	// Call the function to generate a new access token
	_, newAccessToken, exp, err := lib.CreateJwtToken(&models.User{UserID: claims["userId"].(string)})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate access token",
			"success": false,
		})
	}

	// Set the new access token to the user's cookies for future requests
	ctx.Cookie(&fiber.Cookie{
		Name:        "access_token",
		Value:       newAccessToken,
		Expires:     time.Now().Add(time.Hour * 1),
		Secure:      true,
		SessionOnly: true,
	})

	return ctx.JSON(fiber.Map{
		"success":       true,
		"exp":           exp,
		"access_token":  newAccessToken,
	})
}

func ListUsers(ctx *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Find(&users)
	return ctx.Status(fiber.StatusOK).JSON(users)
}

func DeleteAllUsers(ctx *fiber.Ctx) error {
	// Perform the deletion
	if err := database.DB.Exec("DELETE FROM users WHERE role = 'ADMIN'").Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "All users deleted",
	})
}
