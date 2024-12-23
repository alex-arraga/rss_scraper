package v1

import (
	"net/http"

	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/alex-arraga/rss_project/internal/api/middlewares"
	"github.com/go-chi/chi"
)

func ProtectedFeedFollowsRoutes(r chi.Router, h handlers.HandlerConfig, authMid func(middlewares.AuthedHandler) http.HandlerFunc) {
	r.Post("/feed_follows", authMid(h.HandlerCreateFeedFollow))
	r.Get("/feed_follows", authMid(h.HandlerGetFeedsFollows))
	r.Delete("/feed_follows/{feedFollowID}", authMid(h.HandlerDeleteFeedFollows))
}
