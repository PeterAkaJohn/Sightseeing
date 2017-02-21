package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/PeterAkaJohn/SightSeeing/src/model"
	"github.com/PeterAkaJohn/SightSeeing/src/viewmodel"
	"github.com/gorilla/mux"
)

type favoriteController struct {
	template *template.Template
}

func (fc *favoriteController) GetUserFavorites(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	userID, err := model.GetUserID(username)
	fmt.Println(userID)
	if err != nil {
		log.Print(err)
	}
	locations, err := model.GetUserFavorites(userID)
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
	session, err := Store.Get(r, "user-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	val := session.Values["userID"].(int)
	userID := int64(val)

	err = json.NewDecoder(r.Body).Decode(&locationVM)
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
	err = model.AddNewUserFavorite(userID, locationID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
}
