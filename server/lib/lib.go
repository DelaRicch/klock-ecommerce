package lib

import (
	"context"
	"fmt"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/DelaRicch/klock-ecommerce/server/models"
	"github.com/alexedwards/argon2id"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateRandomStrings(length int) string {
	const charset = "01234ABCDEFGHIJKLM567890NOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}


// Generate Unique User IDs
func GenerateID(value string) string {
	randomString := GenerateRandomStrings(15)
	generatedID := fmt.Sprintf("%s-%s", value,randomString)
	return generatedID
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

// Create new JWT access token and refresh token
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

// Set cloudinary credentials
func Credentials() (*cloudinary.Cloudinary, error) {

    // cld.Config.URL.Secure = true
    cldName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cldSecret := os.Getenv("CLOUDINARY_API_SECRET")
    cldKey := os.Getenv("CLOUDINARY_API_KEY")


    cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
    if err != nil {
        return nil, err
    }
    return cld, nil
}

func UploadToCloudinary(file *multipart.FileHeader, subFolder, productID string) (string, error) {
    ctx := context.Background()
    cld, err := Credentials()
    if err != nil {
        return "", err
    }

    // Open the uploaded file
    openedFile, err := file.Open()
    if err != nil {
        return "", err
    }
    defer openedFile.Close()

    // Construct the folder path in Cloudinary
    folderPath := fmt.Sprintf("klock-ecommerce/%s/%s", subFolder, productID)

	// Remove file extension from file name
	fileName := filepath.Base(file.Filename)
	nameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]

    uploadParams := uploader.UploadParams{
        PublicID: folderPath + "/" + nameWithoutExt,
    }

    result, err := cld.Upload.Upload(ctx, openedFile, uploadParams)
    if err != nil {
        return "", err
    }

    imageUrl := result.SecureURL
    return imageUrl, nil
}