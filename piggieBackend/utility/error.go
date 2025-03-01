package utility

import "errors"

// UserCredentials validation errors
// Used for registration and logging in
var (
	ErrInvalidTextContent = errors.New("error: Invalid credential content")
	ErrInvalidRegexp      = errors.New("error: Invalid regexp formula")
)

// User credentials validation errors
// Used for registration
var (
	ErrInvalidLength            = errors.New("error: Invalid credential length")
	ErrSyntaxRequirementsNotMet = errors.New("error: Syntax requirements not met")
)

// User credentials validation errors
// Used for logging in
var (
	ErrEmptyCredential = errors.New("error: Empty credential form")
)

// Used both for logging in and registration
var (
	// Compared password are not equal
	ErrPasswordMismatch = errors.New("error: Password missmatch")

	// Internal database server error
	ErrDatabaseError = errors.New("error: Something went wrong wile communicating with database server")

	// No rows were returned when they were expected
	ErrNoRows = errors.New("error: No rows were returned")
)

// Used for users session
var (
	ErrInvalidJWT = errors.New("error: JWT that comes with request is invalid")
)

// Get handles helper errors
var (
	ErrInvalidMethod = errors.New("error: Invalid request method. Only GET request are valid")
)
