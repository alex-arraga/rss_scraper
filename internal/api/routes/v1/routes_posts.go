package v1

import (
	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/go-chi/chi"
)

func ProtectedPostsRoutes(r chi.Router, h handlers.HandlerConfig) {
	// r.Get("/posts", apiCfg.MiddlewareAuth(apiCfg.HandlerGetPostsForUser))
}
