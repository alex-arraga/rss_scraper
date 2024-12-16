package routes

import (
	"github.com/alex-arraga/rss_project/internal/api"
	"github.com/go-chi/chi"
)

func PostsRoutes(r chi.Router, apiCfg *api.APIConfig) {
	r.Get("/posts", apiCfg.MiddlewareAuth(apiCfg.HandlerGetPostsForUser))
}
