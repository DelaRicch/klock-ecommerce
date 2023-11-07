package lib

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"

	"golang.org/x/crypto/argon2"
)

// Generate Unique User IDs
func GenerateUserID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand = rand.Reader
	var length = 15

	randomBytes := make([]byte, length)
	for i := range randomBytes {
		max := big.NewInt(int64(len(charset)))
		index, err := rand.Int(seededRand, max)
		if err != nil {
			panic(err)
		}
		randomBytes[i] = charset[index.Int64()]
	}
	return string(randomBytes)
}

// Hash the user's password using Argon2
func HashPassword(password string) (string, error) {
    salt := make([]byte, 16)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }
    hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

    // Encode the salt and hash as base64 and store it
    return base64.StdEncoding.EncodeToString(salt) + "$" + base64.StdEncoding.EncodeToString(hash), nil
}