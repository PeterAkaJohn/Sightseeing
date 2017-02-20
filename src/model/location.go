package model

import "log"

//Location : when the user favorites a location this gets inserted in the database
type Location struct {
	ID          int64  `json:"id"`
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
func GetLocation(locationID int64) (*Location, error) {
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

//GetLocations : retrieves all the locations in the database
func GetLocations() ([]*Location, error) {
	result := []*Location{}

	rows, err := db.Query("SELECT id, name, position, address, city, state, postal_code, description, open_hours, close_hours " +
		"FROM location")
	if err != nil {
		log.Print(err)
	} else {
		for rows.Next() {
			location := Location{}
			rows.Scan(&location.ID, &location.Name, &location.Position, &location.Address, &location.City, &location.State,
				&location.PostalCode, &location.Description, &location.OpenHours, &location.CloseHours)
			result = append(result, &location)
		}
	}

	return result, err
}

//AddLocation : called when favoriting a location
func AddLocation(name string, position string, address string, city string, state string, postalCode string, description string, openHours string, closeHours string) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO location(name, position, address, city, state, postal_code, description, open_hours, close_hours) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id")
	if err != nil {
		log.Print(err)
	}
	var lastID int64
	err = stmt.QueryRow(name, position, address, city, state, postalCode, description, openHours, closeHours).Scan(&lastID)
	if err != nil {
		log.Print(err)
	}
	log.Printf("ID = %d\n", lastID)

	return lastID, err
}
