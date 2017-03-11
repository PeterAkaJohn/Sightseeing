package model

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

//User : can log in through service providers or through local authentication
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

//Provider are used to populate the Providers in the User when creating a new profile
type Provider struct {
	ID          int
	UserID      int
	Name        string
	AccessToken string
}

//Login allows the user to save favorites
func Login(username string, password string) (*User, error) {
	result := User{}

	row := db.QueryRow("SELECT id, username, firstname, lastname, email, password FROM users WHERE username = $1", username)

	err := row.Scan(&result.ID, &result.Username, &result.FirstName, &result.LastName, &result.Email, &result.Password)

	if err != nil {
		log.Print(err)
		return &User{}, errors.New("User doesn't exists or password is invalid")
	}

	err = checkUserPassword(result.Password, password)
	if err != nil {
		return &User{}, errors.New("User doesn't exists or password is invalid")
	}
	return &result, err
}

func checkUserPassword(storedHashedPassword string, passwordFromUser string) error {
	err := bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(passwordFromUser))
	return err
}

//GetUserID : retrieves userID with the username
func GetUserID(username string) (int64, error) {
	var userID int64

	row := db.QueryRow("SELECT id FROM users WHERE username = $1", username)

	err := row.Scan(&userID)

	if err != nil || userID == 0 {
		log.Print(err)
		return 0, errors.New("User Not Found")
	}

	return userID, err
}

//Register the users
func Register(username string, passwordToCheck string, firstname string, lastname string, email string) error {
	password := []byte(passwordToCheck)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		return err
	}

	stmt, err := db.Prepare("INSERT INTO users(username, firstname, lastname, email, password, created_at) VALUES($1, $2, $3, $4, $5, DEFAULT) RETURNING id")
	if err != nil {
		log.Print(err)
	}
	res, err := stmt.Exec(username, firstname, lastname, email, hashedPassword)
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
