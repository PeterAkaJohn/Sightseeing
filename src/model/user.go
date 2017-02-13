package model

import "log"

//User : can log in through service providers or through local authentication
type User struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

//Provider are used to populate the Providers in the User when creating a new profile
type Provider struct {
	ID          int
	UserID      int
	Name        string
	AccessToken string
}

//Login logs the user in the website
func Login(username string, password string) {

}

//Register the users
func Register(username string, password string, firstname string, lastname string, email string) {

}

//GetUser retrieves a user
func GetUser(userID int) (*User, error) {
	result := User{}

	row := db.QueryRow("SELECT id, username, firstname, lastname, email, password FROM users WHERE id = $1", userID)

	err := row.Scan(&result.ID, &result.Username, &result.FirstName, &result.LastName, &result.Email, &result.Password)

	return &result, err

}

//GetProviders Retrieves the providers of the user
func GetProviders(userID int) ([]*Provider, error) {
	result := []*Provider{}

	rows, err := db.Query("SELECT id, user_id, name, access_token FROM providers WHERE id = $1", userID)

	if err != nil {
		log.Print(err)
	} else {
		for rows.Next() {
			provider := Provider{}
			rows.Scan(&provider.ID, &provider.UserID, &provider.Name, &provider.AccessToken)
			result = append(result, &provider)
		}
	}

	return result, err

}
