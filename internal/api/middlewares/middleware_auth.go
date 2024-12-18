package middlewares

import (
	"fmt"
	"net/http"

	"github.com/alex-arraga/rss_project/internal/auth"
	"github.com/alex-arraga/rss_project/internal/utils"
)

type MiddlewareConfig struct {
	AuthService *auth.AuthService
}

// MiddlewareAuth auth the user and set user data in the context
func (m *MiddlewareConfig) MiddlewareAuth(next http.Handler) http.Handler {
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

		// Add user in the context
		ctx := auth.AddUserToContext(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
