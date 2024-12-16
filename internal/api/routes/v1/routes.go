package v1

import (
	"github.com/alex-arraga/rss_project/internal/api"
	"github.com/go-chi/chi"
)

func RegisterV1Routes(r chi.Router, apiCfg *api.APIConfig) {
	v1Router := chi.NewRouter()
	r.Mount("/v1", v1Router)

	UsersRoutes(v1Router, apiCfg)
	FeedsRoutes(v1Router, apiCfg)
	FeedFollowsRoutes(v1Router, apiCfg)
	PostsRoutes(v1Router, apiCfg)
}
