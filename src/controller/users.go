package controller

import (
	"encoding/json"
	"net/http"

	"github.com/PeterAkaJohn/SightSeeing/src/model"
)

type userController struct {
}

func (uc *userController) Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == POST {
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		firstname := r.Form.Get("firstname")
		lastname := r.Form.Get("lastname")
		email := r.Form.Get("email")

		err := model.Register(username, password, firstname, lastname, email)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}
	}
}

func (uc *userController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == POST {
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		user, err := model.Login(username, password)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
