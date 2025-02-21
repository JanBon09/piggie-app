package security

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"piggieBackend/content"

	"golang.org/x/crypto/argon2"
)

// Struct for holding argon2 parameters
type argon2Params struct {
	time      uint32
	memory    uint32
	threads   uint8
	keyLength uint32
}

var hashingParams argon2Params = argon2Params{3, 64 * 1024, 2, 32}

// Generating salt for password hashing
func generateSalt(len uint16) string {
	salt := make([]byte, len)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal("Something went wrong")
	}

	return encodeBytesArray(salt)
}

// Encoding bytes array into string
func encodeBytesArray(arr []byte) string {
	return base64.RawStdEncoding.EncodeToString(arr)
}

// Decoding string into bytes array
func decodeString(text string) ([]byte, error) {
	arr, err := base64.RawStdEncoding.DecodeString(text)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// Generating salt and hashing password using argon2
func hashPasswordNewUser(newUser *content.NewUser, saltLen uint16) (err error) {
	decodedPassword := []byte(newUser.Password)
	salt := generateSalt(saltLen)

	hashedPassword := argon2.IDKey(decodedPassword, []byte(salt), hashingParams.time, hashingParams.memory, hashingParams.threads, hashingParams.keyLength)

	newUser.Password = encodeBytesArray(hashedPassword)
	newUser.Salt = salt

	return nil
}

func hashPasswordExistingUser(existingUser *content.ExistingUser) (err error) {
	decodedPassword := []byte(existingUser.Password)
	decodedSalt := []byte(existingUser.Salt)

	hashedPassword := argon2.IDKey(decodedPassword, decodedSalt, hashingParams.time, hashingParams.memory, hashingParams.threads, hashingParams.keyLength)

	existingUser.Password = encodeBytesArray(hashedPassword)

	return nil
}
