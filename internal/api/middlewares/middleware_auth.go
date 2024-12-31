package middlewares

import (
	"fmt"
	"net/http"

	"github.com/alex-arraga/rss_project/internal/auth"
	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/utils"
)

type MiddlewareConfig struct {
	AuthService *auth.AuthService
}

type AuthedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

// MiddlewareAuth auth the user and set user data in the context
func (m *MiddlewareConfig) MiddlewareAuth(handler AuthedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract ApiKey from header
		apiKey, err := auth.ExtractAPIKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusForbidden, fmt.Sprintf("Auth error: %v", err))
			return
		}

		// Auth user
		user, err := m.AuthService.AuthenticateUser(r.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		// Gives control to the handler, injecting to the user
		handler(w, r, user)
	})
}
