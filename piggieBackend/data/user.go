package data

type NewUser struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	Email             string `json:"email"`
	DateOfBirth       string `json:"dateOfBirth`
	Salt              string `json:"salt"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Country           int16  `json:"country"`
	ProfilePictureURL string `json:"profilePictureURL"`
}
