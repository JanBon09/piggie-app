package server

import (
	"encoding/json"
	"net/http"
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
// To do:
// - Password hashing
// - Optional data handling
func handleRegister(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var registeringUser newUserRequired
	if err := json.NewDecoder(request.Body).Decode(&registeringUser); err != nil {
		http.Error(writer, "Bad request", http.StatusBadRequest)
		return
	}

}
