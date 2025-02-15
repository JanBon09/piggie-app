package security

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"piggieBackend/content"
	"piggieBackend/data"
)

// Generating salt for password hashing
func generateSalt(len int16) string {
	salt := make([]byte, len)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal("Something went wrong")
	}

	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)

	return encodedSalt
}

func ValidateNewUserData(newUser content.NewUser) {
	data.RegisterNewUserRequired(newUser)
}
