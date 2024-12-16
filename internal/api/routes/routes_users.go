package routes

import (
	"github.com/alex-arraga/rss_project/internal/api"
	"github.com/go-chi/chi"
)

func UsersRoutes(r chi.Router, apiCfg *api.APIConfig) {
	r.Post("/users", apiCfg.HandlerCreateUser) //createUser
	r.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUserByAPIKey))
}