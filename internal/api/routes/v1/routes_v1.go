package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/alex-arraga/rss_project/internal/api/middlewares"
)

func RegisterProtectedV1Routes(r chi.Router, h handlers.HandlerConfig, authMid func(middlewares.AuthedHandler) http.HandlerFunc) {
	ProtectedUserRoutes(r, h, authMid)
	ProtectedFeedFollowsRoutes(r, h, authMid)
	ProtectedFeedsRoutes(r, h, authMid)
	ProtectedPostsRoutes(r, h, authMid)
}

func RegisterPublicV1Routes(r chi.Router, h handlers.HandlerConfig) {
	PublicUsersRoutes(r, h)
	PublicFeedsRoutes(r, h)
}
