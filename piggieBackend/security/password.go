package security

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateSalt(len int16) string {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal("Something went wrong")
	}

	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)

	return encodedSalt
}
