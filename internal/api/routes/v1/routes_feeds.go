package v1

import (
	"github.com/alex-arraga/rss_project/internal/api"
	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/alex-arraga/rss_project/internal/services"
	"github.com/go-chi/chi"
)

func FeedsRoutes(r chi.Router, apiCfg *api.APIConfig, srv *services.ServicesConfig) {
	handlerConfig := handlers.HandlerConfig{
		Services: srv,
	}

	r.Get("/feeds", handlerConfig.HandlerGetFeeds)
	// r.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
	// r.Put("/feeds/{feedID}", apiCfg.MiddlewareAuth(apiCfg.HandlerUpdateFeed))
	// r.Delete("/feeds/{feedID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeed))
}
