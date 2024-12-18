package middlewares

import (
	"fmt"
	"net/http"

	"github.com/alex-arraga/rss_project/internal/auth"
	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/utils"
)

type MiddlewareConfig struct {
	DB *database.Queries
}

// type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (m *MiddlewareConfig) MiddlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.ExtractAPIKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusForbidden, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := m.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		ctx := r.Context()
		ctx = AddUserToContext(ctx, user)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
