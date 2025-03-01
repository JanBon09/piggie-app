package security

import (
	"net/http"
	"piggieBackend/data"
	"piggieBackend/utility"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var encodedSecretKey string

var (
	getFunctionsMap map[string]interface{}
)

func initFunctionsMap() {
	getFunctionsMap["userPanel"] = data.GetMainPanelWalletData
}

func LoadSecretKey(filepath string) error {
	var err error

	encodedSecretKey, err = utility.ReadFile(filepath)
	if err != nil {
		return err
	}

	return nil
}

// Create claims for JWT
func claimsCreation(username string, duration int) jwt.MapClaims {
	return jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Duration(duration) * time.Minute).Unix()}
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

// Template function to fill jwt.Parse() method
func retriveSecretKey(token *jwt.Token) (interface{}, error) {
	return encodedSecretKey, nil
}

func parseJWT(userJWTString string) (*jwt.Token, error) {
	return jwt.Parse(userJWTString, retriveSecretKey)
}

// Function takes signed JWT string, parses it and returns username for
// further use in acquiring data
func retriveJWTSubject(userJWT *jwt.Token) (string, error) {
	return userJWT.Claims.GetSubject()
}

// Verifing JWT that comes with user request is valid
func verifyJWT(userJWT *jwt.Token) error {
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
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	return sessionCookie, nil
}

// Function that connects smaller security functions being runed
// for JWT being sent with user's data request.
// 1. It parses signed JWT string and returns it in form of a jwt.Token
// 2. It checks if JWT coming with request is valid
// 3. It retrives username from JWT
// Afther that it goes down the server pipeline to data layer
func UserSessionVerification(userJWTString string) (string, error) {
	userJWT, err := parseJWT(userJWTString)
	if err != nil {
		return "", err
	}

	err = verifyJWT(userJWT)
	if err != nil {
		return "", err
	}

	username, err := retriveJWTSubject(userJWT)
	if err != nil {
		return "", err
	}

	return username, nil
}
