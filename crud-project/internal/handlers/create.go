package handlers

import (
	"net/http"
	"text/template"
)

func MakeTemplate(path string) *template.Template {
	files := []string{path, "static/base.html"}
	return template.Must(template.ParseFiles(files...))
}

var (
	HomeTmpl    = MakeTemplate("static/index.html")
	CreateTmpl  = MakeTemplate("static/create.html")
	ViewTmpl    = MakeTemplate("static/view.html")
	UpdateTmpl  = MakeTemplate("static/update.html")
	DeleteTmpl  = MakeTemplate("static/delete.html")
	UnknownTmpl = MakeTemplate("static/unknown.html")
)

func Create(w http.ResponseWriter, r *http.Request) {
	err := CreateTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
