package api

import (
	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
)

type APIConfig struct {
	DB *database.Queries
}
