package handlers

import (
	"net/http"
)

func View(w http.ResponseWriter, r *http.Request) {
	err := ViewTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
