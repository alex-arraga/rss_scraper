package services

import (
	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
)

type ServicesConfig struct {
	DB *database.Queries
}
