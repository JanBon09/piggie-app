package security

import (
	"net/http"
	"piggieBackend/utility"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var encodedSecretKey string

func LoadSecretKey(filepath string) error {
	var err error

	encodedSecretKey, err = utility.ReadFile("C:\\WebApp_Projects\\piggieApp\\piggieBackend\\utility\\secretKey.txt")
	if err != nil {
		return err
	}

	return nil
}

// Create claims for JWT
func claimsCreation(username string, duration int) jwt.MapClaims {
	return jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Duration(duration) * time.Minute).Unix()}
}

// Create JWT
// Used after a successful user verification
// Returns JWT signed string for usage in cookies
func createJWT(username string, duration int) (string, error) {
	if len(encodedSecretKey) < 1 {
		return "", utility.ErrInvalidLength
	}

	duration /= 60

	claims := claimsCreation(username, duration)
	userJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return userJWT.SignedString([]byte(encodedSecretKey))
}

// Verifing JWT that comes with user request is valid
func verifyJWT(userJWTString string) error {
	userJWT, err := jwt.Parse(userJWTString, func(token *jwt.Token) (interface{}, error) { return encodedSecretKey, nil })
	if err != nil {
		return err
	}

	if !userJWT.Valid {
		return utility.ErrInvalidJWT
	}

	return nil
}

// Creating cookie with JWT for securing user requests
func UserSessionCookieCreation(username string, duration int) (http.Cookie, error) {
	userJWTString, err := createJWT(username, duration)
	if err != nil {
		return http.Cookie{
			Name:   "",
			Value:  "",
			MaxAge: -1,
		}, err
	}

	sessionCookie := http.Cookie{
		Name:     "userSession",
		Value:    userJWTString,
		MaxAge:   int(duration),
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	return sessionCookie, nil
}
