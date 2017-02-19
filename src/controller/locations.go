package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/PeterAkaJohn/SightSeeing/src/model"
	"github.com/gorilla/mux"
)

type locationController struct {
}

func (lc *locationController) GetLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	locationID, err := strconv.ParseInt(vars["locationID"], 0, 64)
	if err != nil {
		log.Print(err)
	}
	location, err := model.GetLocation(locationID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)

}

func (lc *locationController) GetLocations(w http.ResponseWriter, r *http.Request) {

	locations, err := model.GetLocations()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)

}
