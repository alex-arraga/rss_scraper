package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/models"
	"github.com/alex-arraga/rss_project/internal/utils"
	"github.com/google/uuid"
)

type UserDatabase interface {
	CreateUser(ctx context.Context, params database.CreateUserParams) (database.User, error)
	GetUserByAPIKey(ctx context.Context, apiKey string) (database.User, error)
	GetPostsForUser(ctx context.Context, params database.GetPostsForUserParams) ([]database.Post, error)
}

type UserService struct {
	DB UserDatabase
}

func (us *UserService) CreateUser(ctx context.Context, name string) (models.User, error) {
	if name == "" {
		return models.User{}, fmt.Errorf("name is required")
	}

	user, err := us.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return models.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return models.ResponseAPIUser(user), nil
}

func (us *UserService) GetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseAPIUser(user))
}

func (us *UserService) GetPostsForUser(ctx context.Context, userID uuid.UUID, limit int32) ([]models.Post, error) {
	posts, err := us.DB.GetPostsForUser(ctx, database.GetPostsForUserParams{
		UserID: userID,
		Limit:  limit,
	})
	if err != nil {
		return []models.Post{}, fmt.Errorf("couldn't get the posts %v ", err)
	}

	return models.ResonseAPIPostsForUser(posts), nil
}
