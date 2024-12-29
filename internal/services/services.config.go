package services

import (
	"context"
	"net/http"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
)

type ServicesConfig struct {
	DB     *database.Queries
	UserDB UserDatabase
}

type UserDatabase interface {
	CreateUser(ctx context.Context, params database.CreateUserParams) (database.User, error)
	GetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User)
	GetPostsForUser(ctx context.Context, params database.GetPostsForUserParams) ([]database.Post, error)
}
