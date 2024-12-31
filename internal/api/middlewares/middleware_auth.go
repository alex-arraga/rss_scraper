package middlewares

import (
	"fmt"
	"net/http"

	"github.com/alex-arraga/rss_project/internal/auth"
	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/utils"
	"github.com/go-chi/chi"
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

type ProtectedRouter struct {
	chi.Router
}

// NewProtectedRouter creates a new router with the authentication middleware applied globally.
func NewProtectedRouter(authMid func(AuthedHandler) http.HandlerFunc) *ProtectedRouter {
	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authMid(func(w http.ResponseWriter, r *http.Request, user database.User) {
				next.ServeHTTP(w, r)
			})(w, r)
		})
	})

	return &ProtectedRouter{Router: r}
}
