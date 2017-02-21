package model

import (
	"errors"
	"log"
)

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
func Login(username string, password string) (*User, error) {
	result := User{}

	row := db.QueryRow("SELECT * FROM users WHERE username = $1", username)

	err := row.Scan(&result.ID, &result.Username, &result.FirstName, &result.LastName, &result.Email, &result.Password)

	if err != nil {
		log.Print(err)
	}

	if result.Password == password {
		return &result, err
	}

	return &User{}, errors.New("User doesn't exists or password is invalid")
}

//GetUserID : retrieves userID with the username
func GetUserID(username string) (int64, error) {
	var userID int64

	row := db.QueryRow("SELECT id FROM users WHERE username = $1", username)

	err := row.Scan(&userID)

	if err != nil {
		log.Print(err)
	}

	return userID, err
}

//Register the users
func Register(username string, password string, firstname string, lastname string, email string) error {
	stmt, err := db.Prepare("INSERT INTO users(username, firstname, lastname, email, password) VALUES($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		log.Print(err)
	}
	res, err := stmt.Exec(username, firstname, lastname, email, password)
	if err != nil {
		log.Print(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Print(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Print(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)

	return err
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
