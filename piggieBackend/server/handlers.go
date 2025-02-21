package server

import (
	"encoding/json"
	"net/http"
	"piggieBackend/content"
	"piggieBackend/security"
	"piggieBackend/utility"
)

// Struct for required user data
type newUserRequired struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	DateOfBirth string `json:"dateOfBirth"`
}

// Struct for optional user data
type newUserOptional struct {
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Country           int8   `json:"country"`
	ProfilePictureURL string `json:"profilePictureURL"`
}

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
}
