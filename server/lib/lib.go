package lib

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/DelaRicch/klock-ecommerce/server/models"
	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateRandomNumbers(length int) string {
	const charset = "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func GenerateRandomAlphabets(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// Generate Unique User IDs
func GenerateUserID() string {
	randomPart1 := GenerateRandomNumbers(5)
	randomPart2 := GenerateRandomAlphabets(5)
	userID := fmt.Sprintf("KLOCK-USER-%s-%s", randomPart1, randomPart2)
	return userID
}

// Hash the user's password using Argon2id
func HashPassword(password string) (string, error) {
	hashedPassword, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}

// Verify the user's password using Argon2id
func VerifyPassword(hashedPassword, password string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, hashedPassword)
	if err != nil {
		return false
	}
	return match
}

func CreateJwtToken(user *models.User) (string, string, int64, error) {
	exp := time.Now().Add(time.Minute * 2).Unix()
	rfExp := time.Now().Add(time.Hour * 24 * 30).Unix()
	claims := jwt.MapClaims{
		"exp":    exp,
		"userId": user.UserID,
		"isAdmin": user.Role == "ADMIN",
	}

	rfClaims := jwt.MapClaims{
		"exp":    rfExp,
		"userId": user.UserID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tkn, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", "", 0, err
	}

	rfToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rfClaims)
	refreshTkn, rfErr := rfToken.SignedString([]byte("rf_secret"))
	if rfErr != nil {
		return "", "", 0, rfErr
	}
	
	return refreshTkn, tkn, exp, nil
	
}
