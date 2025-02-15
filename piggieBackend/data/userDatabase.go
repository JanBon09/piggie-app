package data

import "piggieBackend/content"

// Registers new user in a database making it possible for him to login into WebApp
func RegisterNewUserRequired(newUser content.NewUser) error {
	query := "INSERT INTO users(username, password, email, dateofbirth, salt) "
	query += "VALUES($1, $2, $3, $4, $5)"
	statement, err := DB.Prepare(query)
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
func VerifyUserExistence(user content.ExistingUser) {}
