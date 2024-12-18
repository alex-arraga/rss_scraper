package v1

import (
	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/go-chi/chi"
)

func RegisterProtectedV1Routes(r chi.Router, h handlers.HandlerConfig) {
	// ProtectedUserRoutes(r, srv)
	ProtectedFeedFollowsRoutes(r, h)
	// ProtectedFeedsRoutes(r, srv)
	// ProtectedPostsRoutes(r, srv)
}

func RegisterPublicV1Routes(r chi.Router, h handlers.HandlerConfig) {
	PublicUsersRoutes(r, h)
	PublicFeedsRoutes(r, h)
}
