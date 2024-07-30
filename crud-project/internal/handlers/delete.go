package handlers

import (
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	err := DeleteTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
