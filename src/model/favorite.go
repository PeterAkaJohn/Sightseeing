package model

import "log"

//Favorite defines the relation between the user and the locations they have as favorites
type Favorite struct {
	ID         int
	UserID     int
	LocationID int
}

//Retrieves all the user liked locations
func GetUserFavorites(userID int) ([]*Favorite, error) {
	result := []*Favorite{}

	rows, err := db.Query("SELECT id, user_id, location_id" +
		"FROM favorites")

	if err != nil {
		log.Print(err)
	} else {
		for rows.Next() {
			favorite := Favorite{}
			rows.Scan(&favorite.ID, &favorite.UserID, &favorite.LocationID)
			result = append(result, &favorite)
		}
	}

	return result, err
}
