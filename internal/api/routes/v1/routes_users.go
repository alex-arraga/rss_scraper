package v1

import (
	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/alex-arraga/rss_project/internal/services"
	"github.com/go-chi/chi"
)

func UsersRoutes(r chi.Router, userService *services.ServicesConfig) {
	handlerConfig := handlers.HandlerConfig{
		UserService: userService,
	}

	r.Post("/users", handlerConfig.HandlerCreateUser)
	// r.Get("/users", handlerConfig.MiddlewareAuth(apiCfg.HandlerGetUserByAPIKey))
}
