package v1

import (
	"github.com/alex-arraga/rss_project/internal/api"
	"github.com/alex-arraga/rss_project/internal/services"
	"github.com/go-chi/chi"
)

// ! I have to remove apiCfg later,
// !  when all routes are updated
func RegisterV1Routes(r chi.Router, apiCfg *api.APIConfig, srv *services.ServicesConfig) {
	v1Router := chi.NewRouter()
	r.Mount("/v1", v1Router)

	UsersRoutes(v1Router, srv)
	FeedsRoutes(v1Router, apiCfg, srv)
	FeedFollowsRoutes(v1Router, apiCfg)
	PostsRoutes(v1Router, apiCfg)
}
