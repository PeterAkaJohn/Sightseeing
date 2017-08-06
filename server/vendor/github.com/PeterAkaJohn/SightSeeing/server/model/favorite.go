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

	rows, err := db.Query("SELECT locations.id, locations.name, locations.position, locations.address, locations.city,"+
		"locations.postal_code, locations.description FROM locations, favorites "+
		"WHERE favorites.user_id = $1 AND locations.id = favorites.location_id", userID)
	if err != nil {
		log.Print(err)
	} else {
		for rows.Next() {
			location := Location{}
			rows.Scan(&location.ID, &location.Name, &location.Position, &location.Address, &location.City,
				&location.PostalCode, &location.Description)
			result = append(result, &location)
		}
	}

	return result, err
}

//AddNewUserFavorite : add a new user favorite location
func AddNewUserFavorite(userID int64, locationID int64) error {
	stmt, err := db.Prepare("INSERT INTO favorites(user_id, location_id, created_at) VALUES($1, $2, DEFAULT)")
	if err != nil {
		log.Print(err)
	}
	_, err = stmt.Exec(userID, locationID)
	if err != nil {
		log.Print(err)
	}
	// lastID, err := res.LastInsertId()
	// if err != nil {
	// 	log.Print(err)
	// }
	// rowCnt, err := res.RowsAffected()
	// if err != nil {
	// 	log.Print(err)
	// }
	// log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)

	return err
}
