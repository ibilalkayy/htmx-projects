package main

import (
	"fmt"
	"log"
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

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		err := UnknownTmpl.Execute(w, nil)
		if err != nil {
			log.Fatal("Error executing the unknown template")
		}
		return
	}
	err := HomeTmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("Error executing the home template")
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	err := CreateTmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("Error executing the create template")
	}
}

func View(w http.ResponseWriter, r *http.Request) {
	err := ViewTmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("Error executing the view template")
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	err := UpdateTmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("Error executing the update template")
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	err := DeleteTmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("Error executing the delete template")
	}
}

func main() {
	http.HandleFunc("GET /", Home)
	http.HandleFunc("GET /create", Create)
	http.HandleFunc("GET /view", View)
	http.HandleFunc("GET /update", Update)
	http.HandleFunc("GET /delete", Delete)
	fmt.Println("Started listening on port: 8080")
	http.ListenAndServe(":8080", nil)
}
