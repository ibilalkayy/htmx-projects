package routes

import (
	"net/http"

	"github.com/ibilalkayy/htmx-projects/crud-project/internal/handlers"
)

func Routes() {
	http.HandleFunc("GET /", handlers.Home)
	http.HandleFunc("GET /create", handlers.Create)
	http.HandleFunc("GET /view", handlers.View)
	http.HandleFunc("GET /update", handlers.Update)
	http.HandleFunc("GET /delete", handlers.Delete)
}
