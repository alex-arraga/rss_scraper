package routes

import (
	"github.com/alex-arraga/rss_project/internal/api"
	"github.com/alex-arraga/rss_project/internal/api/handlers"
	v1 "github.com/alex-arraga/rss_project/internal/api/routes/v1"
	"github.com/alex-arraga/rss_project/internal/services"
	"github.com/go-chi/chi"
)

// ! I have to remove apiCfg later,
// !  when all routes are updated
func RegisterRoutes(r chi.Router, apiCfg *api.APIConfig, srv *services.ServicesConfig) {
	handlerConfig := handlers.HandlerConfig{
		Services: srv,
	}

	// General routes
	r.Get("/healthz", handlerConfig.HealthyHandler)
	r.Get("/err", handlerConfig.UnhealthyHandler)

	// V1 Routes
	v1.RegisterV1Routes(r, apiCfg, srv)
}
