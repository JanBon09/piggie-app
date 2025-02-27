package server

import (
	"log"
	"net/http"
)

func allowCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// Seems a little bit overshot, should work in the same fashion as "Access-Control-Allow-Origin": "*"
		writer.Header().Set("Access-Control-Allow-Origin", "*")

		// Allowed methods
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Allowed types of content
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Allowed credentials like cookie and certs
		writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if request.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(writer, request)
	})
}

func InitServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", handleWelcome)
	mux.HandleFunc("/register", handleRegister)
	mux.HandleFunc("/login", handleLogin)

	err := http.ListenAndServe(":8080", allowCORS(mux))
	if err != nil {
		log.Fatal(err)
	}
}
