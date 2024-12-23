package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
)

type AuthService struct {
	DB *database.Queries
}

type ContextKey string

const UserContextKey = ContextKey("user")

// ExtractAPIKey extracts an API Key from
// the headers of an HTTP request
// Example:
// Authorization: ApiKey {insert apikey here}
func ExtractAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authentication data not found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	return vals[1], nil
}

func (as *AuthService) AuthenticateUser(ctx context.Context, apiKey string) (database.User, error) {
	user, err := as.DB.GetUserByAPIKey(ctx, apiKey)
	if err != nil {
		return database.User{}, errors.New("invalid API Key")
	}
	return user, nil
}
