package handlers

import (
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	err := UpdateTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
