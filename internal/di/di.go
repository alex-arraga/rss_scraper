package di

import (
	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/services"
)

type Container struct {
	UserService services.UserService
}

func NewContainer(db *database.Queries) (*Container, error) {
	userService := &services.UserService{DB: db}

	return &Container{
		UserService: *userService,
	}, nil
}
