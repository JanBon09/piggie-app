package content

// Struct representing and legitimizing user in a Piggie WebApp
type NewUser struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	Email             string `json:"email"`
	DateOfBirth       string `json:"dateOfBirth"`
	Salt              string `json:"salt"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Country           int16  `json:"country"`
	ProfilePictureURL string `json:"profilePictureURL"`
}

// Struct representing user that tries to login
type ExistingUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

// Struct used in a process of verifiaction user credentials that tries to login
type PasswordAndSalt struct {
	Password string `json:"password"`
	Salt     string `json:"salt"`
}
