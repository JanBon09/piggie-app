package security

import (
	"errors"
	"piggieBackend/content"
	"regexp"
)

var (
	ErrInvalidLength            = errors.New("error: Invalid credential length")
	ErrSyntaxRequirementsNotMet = errors.New("error: Syntax requirements not met")
	ErrInvalidRegexp            = errors.New("error: Invalid regexp formula")
	ErrInvalidTextContent       = errors.New("error: Invalid credential content")
)

func checkUsername(username string) error {
	if len(username) < 1 {
		return ErrInvalidLength
	}

	r, err := regexp.Compile(`[{}\[\]()]`)
	if err != nil {
		return ErrInvalidRegexp
	}

	if r.MatchString(username) {
		return ErrInvalidTextContent
	}

	return nil
}

func checkPassword(password string) {

}

func CheckEmail(email string) error {
	if len(email) < 1 {
		return ErrInvalidLength
	}

	r, err := regexp.Compile("^[a-zA-Z0-9]{1,}@[a-zA-Z]{1,}.[a-zA-Z]{1,4}$")
	if err != nil {
		return ErrInvalidRegexp
	}

	if !r.MatchString(email) {
		return ErrSyntaxRequirementsNotMet
	}

	return nil
}

func SecurityProcessNewUser(newUser content.NewUser) {

}
