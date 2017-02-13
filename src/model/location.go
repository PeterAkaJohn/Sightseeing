package model

//Location : when the user favorites a location this gets inserted in the database
type Location struct {
	ID          int
	Name        string
	Position    string
	Address     string
	City        string
	State       string
	PostalCode  string
	Description string
	OpenHours   string
	CloseHours  string
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
