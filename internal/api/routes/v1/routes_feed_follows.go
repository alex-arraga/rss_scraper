package v1

import (
	"github.com/alex-arraga/rss_project/internal/api"
	"github.com/go-chi/chi"
)

func FeedFollowsRoutes(r chi.Router, apiCfg *api.APIConfig) {
	r.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	r.Get("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedsFollows))
	r.Delete("/feed_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollows))
}
