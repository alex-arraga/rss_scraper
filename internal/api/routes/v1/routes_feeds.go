package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/alex-arraga/rss_project/internal/api/middlewares"
)

func PublicFeedsRoutes(r chi.Router, h handlers.HandlerConfig) {
	r.Get("/feeds", h.HandlerGetFeeds)
}

func ProtectedFeedsRoutes(r chi.Router, h handlers.HandlerConfig, authMid func(middlewares.AuthedHandler) http.HandlerFunc) {
	r.Post("/feeds", authMid(h.HandlerCreateFeed))
	r.Put("/feeds/{feedID}", authMid(h.HandlerUpdateFeed))
	r.Delete("/feeds/{feedID}", authMid(h.HandlerDeleteFeed))
}
