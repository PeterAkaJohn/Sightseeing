package controller

import (
	"net/http"
	"text/template"
)

type indexController struct {
	Template *template.Template
}

func (lc *indexController) Main(w http.ResponseWriter, r *http.Request) {
	lc.Template.Execute(w, nil)
}
