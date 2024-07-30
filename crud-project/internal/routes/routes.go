package routes

import (
	"net/http"

	"github.com/ibilalkayy/htmx-projects/crud-project/internal/handler"
)

func Routes() {
	tmpl := handler.Handler{}
	http.HandleFunc("GET /", tmpl.Deps.Home)
	http.HandleFunc("GET /create", tmpl.Deps.Create)
	http.HandleFunc("GET /view", tmpl.Deps.View)
	http.HandleFunc("GET /update", tmpl.Deps.Update)
	http.HandleFunc("GET /delete", tmpl.Deps.Delete)
}
