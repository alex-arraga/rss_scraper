package v1

import (
	"github.com/alex-arraga/rss_project/internal/services"
	"github.com/go-chi/chi"
)

func RegisterProtectedV1Routes(r chi.Router, srv *services.ServicesConfig) {
	// ProtectedUserRoutes(r, srv)
	ProtectedFeedFollowsRoutes(r, srv)
	// ProtectedFeedsRoutes(r, srv)
	// ProtectedPostsRoutes(r, srv)
}

func RegisterPublicV1Routes(r chi.Router, srv *services.ServicesConfig) {
	PublicUsersRoutes(r, srv)
	PublicFeedsRoutes(r, srv)
}
