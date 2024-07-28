package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "Film 1", Director: "Director 1"},
			{Title: "Film 2", Director: "Director 2"},
			{Title: "Film 3", Director: "Director 3"},
		},
	}
	err := tmpl.Execute(w, films)
	if err != nil {
		log.Fatal("Error executing the template: ", err)
	}
}

func AddFilm(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/add-film/", AddFilm)
	log.Println("Server starting on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
