package security

import (
	"fmt"
	"piggieBackend/content"
	"piggieBackend/data"
	"piggieBackend/utility"
	"regexp"
)

// Checking if username:
// - is atleast 1 character long
// - does not contain any brackets
func checkUsername(username string) error {
	if len(username) < 1 {
		return utility.ErrInvalidLength
	}

	r, err := regexp.Compile(`[{}\[\]()]`)
	if err != nil {
		return utility.ErrInvalidRegexp
	}

	if r.MatchString(username) {
		return utility.ErrInvalidTextContent
	}

	return nil
}

// Checking if password:
// - is atleast 8 characters long
// - contains atleast 1 small letter
// - contains atleast 1 capital letter
// - contains atleast 1 number
// - contains atleast 1 special character
// - does not contain any brackets
func checkPassword(password string) error {
	if len(password) < 8 {
		return utility.ErrInvalidLength
	}

	patterns := make([]string, 5)
	patterns[0] = `[a-z]{1,}`
	patterns[1] = `[A-Z]{1,}`
	patterns[2] = `[0-9]{1,}`
	patterns[3] = `[!@$%^&*,/?\-_+=]`
	patterns[4] = `[(){}\[\]<>]`

	for i := range patterns {
		r, err := regexp.Compile(patterns[i])
		if err != nil {
			return utility.ErrInvalidRegexp
		}

		if i < 4 {
			if !r.MatchString(password) {
				return utility.ErrSyntaxRequirementsNotMet
			}
		} else {
			if r.MatchString(password) {
				return utility.ErrInvalidTextContent
			}
		}
	}

	return nil
}

// Checking if email:
// - is 5 characters long, a@a.a
// - is formulated like: [a-zA-Z0-9]@[a-zA-Z].[a-zA-Z]
func checkEmail(email string) error {
	if len(email) < 5 {
		return utility.ErrInvalidLength
	}

	r, err := regexp.Compile("^[a-zA-Z0-9]{1,}@[a-zA-Z]{1,}.[a-zA-Z]{1,4}$")
	if err != nil {
		return utility.ErrInvalidRegexp
	}

	if !r.MatchString(email) {
		return utility.ErrSyntaxRequirementsNotMet
	}

	return nil
}

// General function to run other function to check user's credentials syntax
func checkUserCredentials(newUser content.NewUser) error {
	if err := checkUsername(newUser.Username); err != nil {
		returnErr := fmt.Errorf("username: %s", err.Error())
		return returnErr
	}
	if err := checkPassword(newUser.Password); err != nil {
		returnErr := fmt.Errorf("password: %s", err.Error())
		return returnErr
	}
	if err := checkEmail(newUser.Email); err != nil {
		returnErr := fmt.Errorf("email: %s", err.Error())
		return returnErr
	}

	return nil
}

// Security layer function for all new user account to run through
// It checks if credentials are valid and meet security requirements
// After checking it runs function to generate salt and hash password
func SecurityRunNewUser(newUser content.NewUser) error {
	if err := checkUserCredentials(newUser); err != nil {
		return err
	}

	if err := hashPasswordNewUser(&newUser, 32); err != nil {
		return err
	}

	if err := data.RegisterNewUserRequired(newUser); err != nil {
		return err
	}

	return nil
}

// While checking username of a existing username
// We make assumption that the only thing we need to protect form
// Is malicious text content being passed using login form
// If user is trying to access his real account, credentials have
// Already had to pass register related validation
func checkExistingCredential(credential string) error {
	if len(credential) == 0 {
		return utility.ErrEmptyCredential
	}

	r, err := regexp.Compile(`[()<>{}\[\]]`)
	if err != nil {
		return utility.ErrInvalidRegexp
	}

	if r.MatchString(credential) {
		return utility.ErrInvalidTextContent
	}

	return nil
}

// Security funciton to validate and verify existence of a user account
// with given credentials
func SecurityRunExistingUser(existingUser content.ExistingUser) error {
	// Validating given credentials - username and password
	if err := checkExistingCredential(existingUser.Username); err != nil {
		return err
	}
	if err := checkExistingCredential(existingUser.Password); err != nil {
		return err
	}

	// Verifing that user account with given username exists, after that retriving hashed password and salt from db
	passwordAndSalt, err := data.VerifyUserExistence(existingUser.Username)
	if err != nil {
		return err
	}

	existingUser.Salt = passwordAndSalt.Salt

	// Hashing user password
	if err := hashPasswordExistingUser(&existingUser); err != nil {
		return err
	}

	// Comparing password retrived from db with password hashed right now using the same salt
	if passwordAndSalt.Password != existingUser.Password {
		return utility.ErrPasswordMismatch
	}

	return nil
}
