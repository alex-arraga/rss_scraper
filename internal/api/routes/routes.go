package routes

import (
	"net/http"

	"github.com/alex-arraga/rss_project/internal/api/handlers"
	v1 "github.com/alex-arraga/rss_project/internal/api/routes/v1"
	"github.com/alex-arraga/rss_project/internal/services"
	"github.com/go-chi/chi"
)

func RegisterRoutes(r chi.Router, srv *services.ServicesConfig, authMid func(next http.Handler) http.Handler) {
	handlerConfig := handlers.HandlerConfig{
		Services: srv,
	}

	// Main subrouter for /v1
	v1Router := chi.NewRouter()

	// Subrouter with auth
	protectedRouter := chi.NewRouter()
	protectedRouter.Use(authMid)

	// V1 Routes
	v1.RegisterProtectedV1Routes(protectedRouter, handlerConfig)
	v1.RegisterPublicV1Routes(v1Router, handlerConfig)

	v1Router.Mount("/", protectedRouter)
	r.Mount("/v1", v1Router)

	// General routes outside of /v1
	r.Get("/healthz", handlerConfig.HealthyHandler)
	r.Get("/err", handlerConfig.UnhealthyHandler)
}
