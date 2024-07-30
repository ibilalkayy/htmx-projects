package handlers

import (
	"net/http"

	"github.com/ibilalkayy/htmx-projects/crud-project/internal/templates"
)

func Create(w http.ResponseWriter, r *http.Request) {
	err := templates.CreateTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
