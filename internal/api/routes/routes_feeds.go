package routes

import (
	"github.com/alex-arraga/rss_project/internal/api"
	"github.com/go-chi/chi"
)

func FeedsRoutes(r chi.Router, apiCfg *api.APIConfig) {
	r.Get("/feeds", apiCfg.HandlerGetFeeds)
	r.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
	r.Put("/feeds/{feedID}", apiCfg.MiddlewareAuth(apiCfg.HandlerUpdateFeed))
	r.Delete("/feeds/{feedID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeed))
}
