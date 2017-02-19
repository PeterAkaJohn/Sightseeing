package viewmodel

//LocationVM : used by controller to communicate when adding a new location
type LocationVM struct {
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
