package di

import (
	"github.com/alex-arraga/rss_project/internal/auth"
	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/services"
)

type Container struct {
	AuthSerive        *auth.AuthService
	UserService       *services.UserService
	FeedService       *services.FeedService
	FeedFollowService *services.FeedFollowService
}

func NewContainer(db *database.Queries) (*Container, error) {
	authService := &auth.AuthService{DB: db}
	userService := &services.UserService{DB: db}
	feedService := &services.FeedService{DB: db}
	feedFollowService := &services.FeedFollowService{DB: db}

	return &Container{
		AuthSerive:        authService,
		UserService:       userService,
		FeedService:       feedService,
		FeedFollowService: feedFollowService,
	}, nil
}
