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

type ExistingUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 'Constructors' of a NewUser struct
func (newUser *NewUser) InitNewUserRequired(username string, password string, email string, dateOfBirth string) {
	newUser.Username = username
	newUser.Password = password
	newUser.Email = email
	newUser.DateOfBirth = dateOfBirth
	newUser.Name = ""
	newUser.Surname = ""
	newUser.Country = -1
	newUser.ProfilePictureURL = ""
}

func (newUser *NewUser) InitNewUserOptional(name string, surname string, country int16, profilePictureURL string) {
	newUser.Name = name
	newUser.Surname = surname
	newUser.Country = country
	newUser.ProfilePictureURL = profilePictureURL
}

func (newUser *NewUser) InitNewUserFull(username string, password string, email string, dateOfBirth string,
	name string, surname string, country int16, profilePictureURL string) {
	newUser.InitNewUserRequired(username, password, email, dateOfBirth)
	newUser.InitNewUserOptional(name, surname, country, profilePictureURL)
}
