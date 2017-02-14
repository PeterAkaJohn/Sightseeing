package model

import "log"

//Location : when the user favorites a location this gets inserted in the database
type Location struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Position    string `json:"position"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postal_code"`
	Description string `json:"description"`
	OpenHours   string `json:"open_hours"`
	CloseHours  string `json:"close_hours"`
}

//GetLocation retrieves the location from the database with the use of the ID
func GetLocation(locationID int) (*Location, error) {
	result := Location{}

	//Use database connection
	row := db.QueryRow(
		"SELECT id, name, position, address, city, state, postal_code, description, open_hours, close_hours "+
			"FROM location "+
			"WHERE id = $1", locationID)

	err := row.Scan(&result.ID, &result.Name, &result.Position, &result.Address, &result.City, &result.State, &result.PostalCode,
		&result.Description, &result.OpenHours, &result.CloseHours)

	return &result, err
}

//GetUserLocations will retrieve all the user liked location to be used in the personal page
func GetUserLocations(userID int) ([]*Location, error) {
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
