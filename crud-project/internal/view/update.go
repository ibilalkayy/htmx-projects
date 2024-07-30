package view

import (
	"net/http"

	"github.com/ibilalkayy/htmx-projects/crud-project/internal/templates"
)

func (MyTemplate) Update(w http.ResponseWriter, r *http.Request) {
	err := templates.UpdateTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
