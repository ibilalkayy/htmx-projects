package handler

import "github.com/ibilalkayy/htmx-projects/crud-project/internal/interfaces"

type Handler struct {
	Deps interfaces.Dependencies
}

func NewHandler(deps interfaces.Dependencies) *Handler {
	return &Handler{Deps: deps}
}
