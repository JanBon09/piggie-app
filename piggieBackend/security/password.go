package security

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/argon2"
)

type argon2Params struct {
	time      uint32
	memory    uint32
	threads   uint8
	keyLength uint32
}

var hashingParams argon2Params = argon2Params{3, 64 * 1024, 2, 32}

// Generating salt for password hashing
func generateSalt(len uint16) []byte {
	salt := make([]byte, len)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal("Something went wrong")
	}

	return salt
}

func encodeBytesArray(arr []byte) string {
	return base64.RawURLEncoding.EncodeToString(arr)
}

func decodeString(text string) ([]byte, error) {
	arr, err := base64.RawStdEncoding.DecodeString(text)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

func HashPassword(password string, saltLen uint16) (encodedHashedPassword string, salt []byte, err error) {
	decodedPassword, err := decodeString(password)
	if err != nil {
		return "", nil, err
	}
	salt = generateSalt(saltLen)

	hashedPassword := argon2.IDKey(decodedPassword, salt, hashingParams.time, hashingParams.memory, hashingParams.threads, hashingParams.keyLength)

	return encodeBytesArray(hashedPassword), salt, nil
}
