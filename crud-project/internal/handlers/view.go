package handlers

import (
	"net/http"

	"github.com/ibilalkayy/htmx-projects/crud-project/internal/templates"
)

func View(w http.ResponseWriter, r *http.Request) {
	err := templates.ViewTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
