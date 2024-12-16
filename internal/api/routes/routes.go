package routes

import (
	"github.com/alex-arraga/rss_project/internal/api"
	v1 "github.com/alex-arraga/rss_project/internal/api/routes/v1"
	"github.com/go-chi/chi"
)

func RegisterRoutes(r chi.Router, apiCfg *api.APIConfig) {
	// General routes
	r.Get("/healthz", apiCfg.HandlerReadiness)
	r.Get("/err", apiCfg.HandlerErr)

	// V1 Routes
	v1.RegisterV1Routes(r, apiCfg)
}
