package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"text/template"

	"github.com/PeterAkaJohn/SightSeeing/src/controller"
	"github.com/gorilla/handlers"
)

func main() {
	templates := populateTemplates()

	go http.ListenAndServe(":8080", http.HandlerFunc(redirect))

	router := controller.Register(templates)

	http.Handle("/", router)
	// http.HandleFunc("/img/", serveResource)
	// http.HandleFunc("/css/", serveResource)

	http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", handlers.LoggingHandler(os.Stdout, router))
}

func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	host, port, err := net.SplitHostPort(req.Host)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(host)
	log.Print(port)
	target := "https://localhost:8443" + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		http.StatusTemporaryRedirect)
}

// func serveResource(w http.ResponseWriter, req *http.Request) {
// 	path := "./public" + req.URL.Path
// 	var contentType string
// 	if strings.HasSuffix(path, ".css") {
// 		contentType = "text/css"
// 	} else if strings.HasSuffix(path, ".png") {
// 		contentType = "image/png"
// 	} else {
// 		contentType = "text/plain"
// 	}
//
// 	f, err := os.Open(path)
//
// 	if err == nil {
// 		defer f.Close()
// 		w.Header().Add("Content-Type", contentType)
//
// 		br := bufio.NewReader(f)
// 		br.WriteTo(w)
// 	} else {
// 		w.WriteHeader(404)
// 	}
// }

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "../public"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths,
				basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}
