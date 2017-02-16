package controller

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/PeterAkaJohn/SightSeeing/src/model"
	"github.com/gorilla/mux"
)

const POST = "POST"
const GET = "GET"
const DELETE = "DELETE"
const UPDATE = "UPDATE"

func sendJSON(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func testFunctionJSON(w http.ResponseWriter, r *http.Request) {
	locations := []model.Location{
		model.Location{
			Name: "Write presentation",
		},
		model.Location{
			Name: "Host meetup",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}

//Register : registers all the routes with the use of the model and later also the viemodel
func Register(templates *template.Template) {
	router := mux.NewRouter()

	indexController := new(indexController)
	indexController.Template = templates.Lookup("index.html")
	router.HandleFunc("/", indexController.Main)

	locationController := new(locationController)
	router.HandleFunc("/browse", locationController.GetLocations)
	router.HandleFunc("/browse/{locationID:[0-9]+}", locationController.GetLocation)

	favoriteController := new(favoriteController)
	router.HandleFunc("/{username}/favorites", favoriteController.GetUserFavorites)
	router.HandleFunc("/{username}/favorites/add", favoriteController.AddFavorite)

	userController := new(userController)
	router.HandleFunc("/register", userController.Register)
	router.HandleFunc("/login", userController.Login)

	http.Handle("/", router)

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)

}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "../public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
