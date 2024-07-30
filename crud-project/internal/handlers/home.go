package handlers

import (
	"net/http"

	"github.com/ibilalkayy/htmx-projects/crud-project/internal/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		err := templates.UnknownTmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}

	err := templates.HomeTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
