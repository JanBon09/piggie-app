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

var (
	ErrPasswordMismatch = errors.New("error: Password missmatch")
	ErrDatabaseError    = errors.New("error: Something went wrong wile communicating with database server")
	ErrNoRows           = errors.New("error: No rows were returned")
)
