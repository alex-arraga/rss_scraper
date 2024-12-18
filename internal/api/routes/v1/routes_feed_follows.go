package v1

import (
	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/go-chi/chi"
)

func ProtectedFeedFollowsRoutes(r chi.Router, h handlers.HandlerConfig) {
	// r.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	// r.Get("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedsFollows))
	// r.Delete("/feed_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollows))
}
