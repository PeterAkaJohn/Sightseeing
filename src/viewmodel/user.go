package viewmodel

//User : viewmodel used when communicating with the model
type UserRegister struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

//User : viewmodel used when communicating with the model
type UserLogIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
