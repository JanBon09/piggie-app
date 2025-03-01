package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"piggieBackend/content"
	"piggieBackend/data"
	"piggieBackend/security"
	"piggieBackend/utility"
)

// Sending simple text content to frontend in form of json
func handleWelcome(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	message, err := json.Marshal("Hello world")
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(message)
}

// Handle for registering new user
// Realized:
// - Reciving and decoding needed data from user from frontend
// - Salt generation
// - Password hashing
// To do:
// - Optional data handling
func handleRegister(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var registeringUser content.NewUser
	if err := json.NewDecoder(request.Body).Decode(&registeringUser); err != nil {
		http.Error(writer, "Bad request", http.StatusBadRequest)
		return
	}

	if err := security.SecurityRunNewUser(registeringUser); err != nil {
		switch err {
		case (utility.ErrInvalidLength):
			http.Error(writer, "Password needs to be atleast 8 characters long", http.StatusForbidden)
		case (utility.ErrSyntaxRequirementsNotMet):
			errMsg := "Password must contain at least 1: capital letter, lowercase letter, number and special character"
			http.Error(writer, errMsg, http.StatusForbidden)
		case (utility.ErrInvalidTextContent):
			http.Error(writer, "Password must not contain any form of brackets", http.StatusForbidden)
		default:
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

// Handling existing user login
func handleLogin(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var existingUser content.ExistingUser
	if err := json.NewDecoder(request.Body).Decode(&existingUser); err != nil {
		http.Error(writer, "Bad request", http.StatusBadRequest)
		return
	}

	if err := security.SecurityRunExistingUser(existingUser); err != nil {
		switch err {
		case utility.ErrPasswordMismatch:
			http.Error(writer, "Wrong username or password", http.StatusUnauthorized)
			return
		case utility.ErrNoRows:
			http.Error(writer, "No user account with given username", http.StatusNotFound)
			return
		default:
			http.Error(writer, "Interal server error", http.StatusInternalServerError)
			return
		}
	}

	userSessionCookie, err := security.UserSessionCookieCreation(existingUser.Username, 60)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Bad request", http.StatusBadRequest)
		return
	}

	http.SetCookie(writer, &userSessionCookie)
}

// Helper for all GET'ters of user data that needs JWT validation
// and retriving user's credentials from it
func userDataHandleHelper(writer *http.ResponseWriter, request *http.Request) (string, error) {
	if request.Method != http.MethodGet {
		http.Error(*writer, "Bad method", http.StatusMethodNotAllowed)
		return "", utility.ErrInvalidMethod
	}

	// Retriving cookie storing user's JWT
	sessionCookie, err := request.Cookie("userSession")
	if err != nil {
		http.Error(*writer, "Bad request", http.StatusBadRequest)
		return "", err
	}

	// Security section for validating JWT coming with request
	username, err := security.UserSessionVerification(sessionCookie.Value)
	if err != nil {
		http.Error(*writer, "Unauthorized", http.StatusUnauthorized)
		return "", err
	}

	return username, nil
}

func handleUserPanel(writer http.ResponseWriter, request *http.Request) {
	// Using helper for JWT validation and retriving user's credentials from it
	// http.Error is omitted here beacuse it was written to writer in helper.
	username, err := userDataHandleHelper(&writer, request)
	if err != nil {
		return
	}

	// Retriving user wallet data
	wallet, err := data.GetMainPanelWalletData(username)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	walletMarshalized, err := json.Marshal(wallet)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(walletMarshalized)
}
