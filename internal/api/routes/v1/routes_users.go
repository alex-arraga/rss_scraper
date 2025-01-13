package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/alex-arraga/rss_project/internal/api/middlewares"
)

func PublicUsersRoutes(r chi.Router, h handlers.HandlerConfig) {
	r.Post("/users", h.HandlerCreateUser)
}

func ProtectedUserRoutes(r chi.Router, h handlers.HandlerConfig, authMid func(middlewares.AuthedHandler) http.HandlerFunc) {
	r.Get("/users", authMid(h.HandlerGetUserByAPIKey))
}
