package controller

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/PeterAkaJohn/SightSeeing/src/model"
)

type favoriteController struct {
	template *template.Template
}

func (fc *favoriteController) GetUserFavorites(w http.ResponseWriter, r *http.Request) {
	if r.Method == GET {
		locations, err := model.GetUserFavorites(1)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(locations)
	}
}

func (fc *favoriteController) AddFavorite(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == POST {
		name := r.Form.Get("name")
		position := r.Form.Get("position")
		address := r.Form.Get("address")
		city := r.Form.Get("city")
		state := r.Form.Get("state")
		postalCode := r.Form.Get("postal_code")
		description := r.Form.Get("description")
		openHours := r.Form.Get("open_hours")
		closeHours := r.Form.Get("close_hours")
		locationID, err := model.AddLocation(name, position, address, city, state, postalCode, description, openHours, closeHours)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}
		err = model.AddNewUserFavorite(1, locationID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
