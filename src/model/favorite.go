package model

import "log"

//Favorite defines the relation between the user and the locations they have as favorites
type Favorite struct {
	ID         int
	UserID     int
	LocationID int
}

//GetUserFavorites : Retrieves all the user liked locations, NOT GOING TO BE USED
func GetUserFavorites(userID int64) ([]*Location, error) {
	result := []*Location{}

	rows, err := db.Query("SELECT location.id, location.name, location.position, location.address, location.city,"+
		"location.postal_code, location.description, location.open_hours, location.close_hours FROM location, favorite "+
		"WHERE favorite.user_id = $1 AND location.id = favorite.location_id", userID)
	if err != nil {
		log.Print(err)
	} else {
		for rows.Next() {
			location := Location{}
			rows.Scan(&location.ID, &location.Name, &location.Position, &location.Address, &location.City,
				&location.PostalCode, &location.Description, &location.OpenHours, &location.CloseHours)
			result = append(result, &location)
		}
	}

	return result, err
}

//AddNewUserFavorite : add a new user favorite location
func AddNewUserFavorite(userID int64, locationID int64) error {
	stmt, err := db.Prepare("INSERT INTO favorites(user_id, location_id) VALUES(?, ?)")
	if err != nil {
		log.Print(err)
	}
	res, err := stmt.Exec(userID, locationID)
	if err != nil {
		log.Print(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)

	return err
}
