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

	router.HandleFunc("/", testFunctionJSON)
	// homeController := new(HomeController)
	// homeController.Template = templates.Lookup("homeShared.html")
	// homeController.LoginTemplate = templates.Lookup("login.html")
	// homeController.RegisterTemplate = templates.Lookup("register.html")
	// router.HandleFunc("/", homeController.Get)
	// router.HandleFunc("/register", homeController.RegisterAndLogIn)
	// router.HandleFunc("/login", homeController.Login)
	//
	// fieldsController := new(FieldsController)
	// fieldsController.Template = templates.Lookup("fields.html")
	// router.HandleFunc("/fields", fieldsController.Get)
	// router.HandleFunc("/testField", fieldsController.JsonTest)
	//
	// fieldController := new(FieldController)
	// fieldController.Template = templates.Lookup("field.html")
	// router.HandleFunc("/fields/{id}", fieldController.Get)
	//
	// postController := new(PostController)
	// postController.Template = templates.Lookup("post.html")
	// router.HandleFunc("/posts/{id}", postController.Get)
	//
	// userController := new(UserController)
	// userController.Template = templates.Lookup("profile.html")
	// router.HandleFunc("/profile", userController.Handle)
	//
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
