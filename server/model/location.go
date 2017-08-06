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
}

//GetLocation retrieves the location from the database with the use of the ID
func GetLocation(locationID int64) (*Location, error) {
	result := Location{}

	//Use database connection
	row := db.QueryRow(
		"SELECT id, name, position, address, city, state, postal_code, description "+
			"FROM locations "+
			"WHERE id = $1", locationID)

	err := row.Scan(&result.ID, &result.Name, &result.Position, &result.Address, &result.City, &result.State, &result.PostalCode,
		&result.Description)

	return &result, err
}

//GetLocations : retrieves all the locations in the database
func GetLocations() ([]*Location, error) {
	result := []*Location{}

	rows, err := db.Query("SELECT id, name, position, address, city, state, postal_code, description " +
		"FROM locations")
	if err != nil {
		log.Print(err)
	} else {
		for rows.Next() {
			location := Location{}
			rows.Scan(&location.ID, &location.Name, &location.Position, &location.Address, &location.City, &location.State,
				&location.PostalCode, &location.Description)
			result = append(result, &location)
		}
	}

	return result, err
}

//AddLocation : called when favoriting a location
func AddLocation(name string, position string, address string, city string, state string, postalCode int, description string) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO locations(name, position, address, city, state, postal_code, description, created_at) VALUES($1, $2, $3, $4, $5, $6, $7, DEFAULT) RETURNING id")
	if err != nil {
		log.Print(err)
	}
	var lastID int64
	err = stmt.QueryRow(name, position, address, city, state, postalCode, description).Scan(&lastID)
	if err != nil {
		log.Print(err)
	}
	log.Printf("ID = %d\n", lastID)

	return lastID, err
}
