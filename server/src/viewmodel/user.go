package viewmodel

type UserVM struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}
