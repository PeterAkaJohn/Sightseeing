package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/PeterAkaJohn/SightSeeing/src/middleware"
	"github.com/PeterAkaJohn/SightSeeing/src/model"
	"github.com/PeterAkaJohn/SightSeeing/src/viewmodel"
	"github.com/gorilla/mux"
)

type favoriteController struct {
	template *template.Template
}

func (fc *favoriteController) GetUserFavorites(w http.ResponseWriter, r *http.Request) {
	// _, err := middleware.VerifyToken(r)
	// if err != nil {
	// 	log.Print(err)
	// 	log.Print("Token Not Valid")
	// 	http.Error(w, err.Error(), http.StatusUnauthorized)
	// 	return
	// }
	vars := mux.Vars(r)
	username := vars["username"]

	userID, err := model.GetUserID(username)
	fmt.Println(userID)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
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
	// claims, err := middleware.VerifyToken(r)
	// if err != nil {
	// 	log.Print(err)
	// 	http.Error(w, err.Error(), http.StatusUnauthorized)
	// 	return
	// }
	//
	// userIDFromClaims := claims["id"].(float64)
	// userID := int64(userIDFromClaims)
	userID := middleware.UserIDFromContext(r.Context())
	var locationVM viewmodel.LocationVM
	//session, err := Store.Get(r, "loginSession")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// val, ok := session.Values["userID"].(int)
	// if !ok {
	// 	fmt.Println("Logged out or not logged in")
	// 	return
	// }
	// userID := int64(val)
	fmt.Println(userID)
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&locationVM)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	locationID, err := model.AddLocation(locationVM.Name, locationVM.Position, locationVM.Address, locationVM.City, locationVM.State, locationVM.PostalCode, locationVM.Description)
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
	//session.Save(r, w)
}
