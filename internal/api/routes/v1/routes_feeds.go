package v1

import (
	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/go-chi/chi"
)

func PublicFeedsRoutes(r chi.Router, h handlers.HandlerConfig) {
	r.Get("/feeds", h.HandlerGetFeeds)
}

func ProtectedFeedsRoutes(r chi.Router, h handlers.HandlerConfig) {
// 	r.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
// 	r.Put("/feeds/{feedID}", apiCfg.MiddlewareAuth(apiCfg.HandlerUpdateFeed))
// 	r.Delete("/feeds/{feedID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeed))}
}
