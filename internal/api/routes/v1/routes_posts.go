package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/alex-arraga/rss_project/internal/api/middlewares"
)

func ProtectedPostsRoutes(r chi.Router, h handlers.HandlerConfig, authMid func(middlewares.AuthedHandler) http.HandlerFunc) {
	r.Get("/posts", authMid(h.HandlerGetPostsForUser))
}
