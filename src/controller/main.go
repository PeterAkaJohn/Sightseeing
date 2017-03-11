package controller

import (
	"net/http"
	"text/template"

	"github.com/PeterAkaJohn/SightSeeing/src/middleware"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

//Store : session to store user_ids and such
var Store = sessions.NewCookieStore([]byte("something-very-secret"))

//MySigningKey : used for jwt authentication
var MySigningKey = []byte("secrets")

//POST :
const POST = "POST"

//GET :
const GET = "GET"

//DELETE :
const DELETE = "DELETE"

//UPDATE :
const UPDATE = "UPDATE"

func init() {
	Store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
	}
}

//Register : registers all the routes with the use of the model and later also the viemodel
func Register(templates *template.Template) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	indexController := new(indexController)
	indexController.Template = templates.Lookup("index.html")
	router.HandleFunc("/", indexController.Main).Methods(GET)

	locationController := new(locationController)
	router.HandleFunc("/browse", locationController.GetLocations).Methods(GET)
	router.HandleFunc("/browse/{locationID:[0-9]+}", locationController.GetLocation).Methods(GET)

	favoriteController := new(favoriteController)
	//GetUserFavoritesHandler := http.HandlerFunc(favoriteController.GetUserFavorites)
	AddFavoriteHandler := http.HandlerFunc(favoriteController.AddFavorite)
	router.HandleFunc("/{username}/favorites", favoriteController.GetUserFavorites).Methods(GET)
	router.Handle("/{username}/favorites/add", middleware.VerifyMiddleware(AddFavoriteHandler)).Methods(POST)

	userController := new(userController)
	logoutHandler := http.HandlerFunc(userController.Logout)
	router.HandleFunc("/register", userController.Register).Methods(POST)
	router.HandleFunc("/login", userController.Login).Methods(POST)
	router.Handle("/logout", middleware.VerifyMiddleware(logoutHandler)).Methods(POST)

	return router

}
