package data

import (
	"database/sql"
	"piggieBackend/content"
	"piggieBackend/utility"
)

// Registers new user in a database making it possible for him to login into WebApp
func RegisterNewUserRequired(newUser content.NewUser) error {
	stringStatement := "INSERT INTO testUsers(username, password, email, dateofbirth, salt) "
	stringStatement += "VALUES($1, $2, $3, $4, $5)"
	statement, err := DB.Prepare(stringStatement)
	if err != nil {
		return err
	}

	_, err = statement.Exec(newUser.Username, newUser.Password, newUser.Email, newUser.DateOfBirth, newUser.Salt)
	if err != nil {
		return err
	}

	return nil
}

// Verify user existence to log him in
// Function returns error if there is any problem in a process of getting data
// Or there is problem with data itself, nil is returned only on success
// Second(Or more like first) return value is true when processed of
// Acquiring data run successfuly but data does not exist or there is a mismatch in data
// False is returned on problem with process of getting data itself
func VerifyUserExistence(username string) (content.PasswordAndSalt, error) {
	stringStatement := "SELECT password, salt FROM testUsers WHERE username LIKE $1"
	statement, err := DB.Prepare(stringStatement)
	if err != nil {
		return content.PasswordAndSalt{Password: "", Salt: ""}, err
	}

	var passwordAndSalt content.PasswordAndSalt
	err = statement.QueryRow(username).Scan(&passwordAndSalt.Password, &passwordAndSalt.Salt)
	if err != nil {
		if err == sql.ErrNoRows {
			return content.PasswordAndSalt{Password: "", Salt: ""}, utility.ErrNoRows
		} else {
			return content.PasswordAndSalt{Password: "", Salt: ""}, err
		}
	}

	return passwordAndSalt, nil
}
