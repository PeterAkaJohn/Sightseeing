package controller

import (
	"encoding/json"
	"net/http"

	"github.com/PeterAkaJohn/SightSeeing/src/model"
	"github.com/PeterAkaJohn/SightSeeing/src/viewmodel"
)

type userController struct {
}

func (uc *userController) Register(w http.ResponseWriter, r *http.Request) {
	var userVM viewmodel.UserRegister
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&userVM)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = model.Register(userVM.Username, userVM.Password, userVM.FirstName, userVM.LastName, userVM.Email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
}

func (uc *userController) Login(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "user-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var userVM viewmodel.UserLogIn
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&userVM)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	user, err := model.Login(userVM.Username, userVM.Password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	session.Values["userID"] = user.ID
	session.Save(r, w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
