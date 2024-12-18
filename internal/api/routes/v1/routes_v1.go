package v1

import (
	"github.com/alex-arraga/rss_project/internal/services"
	"github.com/go-chi/chi"
)

func RegisterProtectedV1Routes(r chi.Router, srv *services.ServicesConfig) {
	v1Router := chi.NewRouter()
	r.Mount("/v1", v1Router)

	ProtectedUserRoutes(v1Router, srv)
	ProtectedFeedFollowsRoutes(v1Router, srv)
	ProtectedFeedsRoutes(v1Router, srv)
	ProtectedPostsRoutes(v1Router, srv)
}

func RegisterPublicV1Routes(r chi.Router, srv *services.ServicesConfig) {
	v1Router := chi.NewRouter()
	r.Mount("/v1", v1Router)

	PublicUsersRoutes(v1Router, srv)
	FeedsRoutes(v1Router, srv)
	FeedFollowsRoutes(v1Router, srv)
	PostsRoutes(v1Router, srv)
}
