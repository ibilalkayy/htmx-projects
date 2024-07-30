package interfaces

import "net/http"

type Templates interface {
	Home(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	View(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type Dependencies struct {
	Templates
}
