package v1

import (
	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/go-chi/chi"
)

func PublicUsersRoutes(r chi.Router, h handlers.HandlerConfig) {
	r.Post("/users", h.HandlerCreateUser)
}

func ProtectedUserRoutes(r chi.Router, h handlers.HandlerConfig) {
	// r.Get("/users", h.HandlerGetUserByAPIKey())
}
