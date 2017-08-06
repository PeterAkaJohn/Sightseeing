package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/PeterAkaJohn/SightSeeing/server/model"
	jwt "github.com/dgrijalva/jwt-go"
)

type userController struct {
}

type tokenResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

func (uc *userController) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = model.Register(user.Username, user.Password, user.FirstName, user.LastName, user.Email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
}

func (uc *userController) Login(w http.ResponseWriter, r *http.Request) {
	//session, err := Store.Get(r, "loginSession")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	var userLog model.User
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&userLog)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	user, err := model.Login(userLog.Username, userLog.Password)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 400)
		return
	}
	//Set jwt token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString(MySigningKey)
	tokenJSON := tokenResponse{
		Token:    tokenString,
		Username: user.Username,
	}
	json.NewEncoder(w).Encode(tokenJSON)
	//userVM will only be used in profile page, using it now only for debugging purposes
	// userVM := converters.ConvertUserToUserVM(*user)
	// session.Values["userID"] = user.ID
	// session.Save(r, w)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(userVM)
}

func (uc *userController) Logout(w http.ResponseWriter, r *http.Request) {
	//More Client Side than server side, remove auth token from localStorage and from the header
	// session, err := Store.Get(r, "loginSession")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	//
	// session.Options = &sessions.Options{
	// 	MaxAge: -1,
	// }
	//
	// session.Save(r, w)
	w.Header().Set("Authorization", "")
}
