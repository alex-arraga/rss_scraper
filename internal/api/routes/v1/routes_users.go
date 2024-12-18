package v1

import (
	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/alex-arraga/rss_project/internal/services"
	"github.com/go-chi/chi"
)

func PublicUsersRoutes(r chi.Router, srv *services.ServicesConfig) {
	handlerConfig := handlers.HandlerConfig{
		Services: srv,
	}

	r.Post("/users", handlerConfig.HandlerCreateUser)
}

func ProtectedUserRoutes(r chi.Router, srv *services.ServicesConfig) {
	handlerConfig := handlers.HandlerConfig{
		Services: srv,
	}

	// ! Original -> r.Get("/users", handlerConfig.MiddlewareAuth(apiCfg.HandlerGetUserByAPIKey))

	r.Get("/users", handlerConfig.HandlerGetUserByAPIKey())
}
