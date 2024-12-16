package routes

import (
	"github.com/alex-arraga/rss_project/internal/api"
	"github.com/go-chi/chi"
)

func UnhealthyRoute(r chi.Router, apiCfg *api.APIConfig) {
	r.Get("/err", apiCfg.HandlerErr)
}
