package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/DelaRicch/klock-ecommerce/klock-backend/database"
	"github.com/DelaRicch/klock-ecommerce/klock-backend/lib"
	"github.com/DelaRicch/klock-ecommerce/klock-backend/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const errorRfTokenMsg string = "Error generating refresh token"
const invalidEmailOrPass string = "Invalid email or password"

func Home(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to Klock E-commerce backend API")
}

func Register(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}

	// Validate input using the validator library
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}

	// Check for the uniqueness of the email
	var existingUser models.User
	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Email already exist",
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
	user.UserID = lib.GenerateID("KLOCK-USER")

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
	result := database.DB.Where("email = ?", loginRequest.Email).First(&user)
	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": invalidEmailOrPass,
		})
	}

	// Verify the password using Argon2id
	if !lib.VerifyPassword(user.Password, loginRequest.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": invalidEmailOrPass,
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
	user.UserID = lib.GenerateID("KLOCK-USER")

	// Set the default role to "USER" if not specified
	if user.Role == "" {
		user.Role = "USER"
	}

	// Add the user to the DB
	result = database.DB.Create(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
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
		"success":      true,
		"exp":          exp,
		"access_token": newAccessToken,
	})
}

func GetUserProfile(ctx *fiber.Ctx) error {
	user, err := lib.ValidateAccessToken(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid access token",
		})
	}

	// Extract user profile from the user object
	userProfile := models.UserProfile{
		Name:     user.Name,
		Email:    user.Email,
		UserID:   user.UserID,
		Role:     user.Role,
		Photo:    user.Photo,
		Phone:    user.Phone,
		Location: user.Location,
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"user":    userProfile,
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	payload := new(models.User)
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}

	res, err := lib.ValidateAccessToken(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid access token",
		})
	}

	result := database.DB.Model(&models.User{}).Where("user_id = ?", res.UserID).Updates(&payload)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error updating user",
		})
	}

	// Retrieve user from the database
	var user models.User
	userResult := database.DB.Where("user_id = ?", res.UserID).First(&user)
	if userResult.RowsAffected == 0 {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	// Extract user profile from the user object
	userProfile := models.UserProfile{
		Name:     user.Name,
		Email:    user.Email,
		UserID:   user.UserID,
		Role:     user.Role,
		Photo:    user.Photo,
		Phone:    user.Phone,
		Location: user.Location,
	}

	fmt.Println(userProfile)

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "User updated successfully",
		"user":    userProfile,
	})

}

func ListUsers(ctx *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Find(&users)

	var usersResponse []models.UserProfile

	for _, user := range users {
		userResponse := models.UserProfile{
			UserID:   user.UserID,
			Name:     user.Name,
			Email:    user.Email,
			Role:     user.Role,
			Photo:    user.Photo,
			Phone:    user.Phone,
			Location: user.Location,
		}
		usersResponse = append(usersResponse, userResponse)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"users":   usersResponse,
	})
}

func DeleteAllUsers(ctx *fiber.Ctx) error {
	if err := database.DB.Exec("DELETE FROM users WHERE role = 'USER'").Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "All users deleted",
	})
}
