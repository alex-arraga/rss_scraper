package middlewares

import (
	"context"
	"errors"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
)

type contextKey string

const userContextKey = contextKey("user")

func AddUserToContext(ctx context.Context, user database.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func GetUserFromContext(ctx context.Context) (database.User, error) {
	user, ok := ctx.Value(userContextKey).(database.User)
	if !ok {
		return database.User{}, errors.New("user not found in context")
	}
	return user, nil
}
