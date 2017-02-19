package controller

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/PeterAkaJohn/SightSeeing/src/model"
	"github.com/PeterAkaJohn/SightSeeing/src/viewmodel"
)

type favoriteController struct {
	template *template.Template
}

func (fc *favoriteController) GetUserFavorites(w http.ResponseWriter, r *http.Request) {
	locations, err := model.GetUserFavorites(1)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)

}

func (fc *favoriteController) AddFavorite(w http.ResponseWriter, r *http.Request) {
	var locationVM viewmodel.LocationVM
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&locationVM)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	locationID, err := model.AddLocation(locationVM.Name, locationVM.Position, locationVM.Address, locationVM.City, locationVM.State, locationVM.PostalCode, locationVM.Description, locationVM.OpenHours, locationVM.CloseHours)
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
}
