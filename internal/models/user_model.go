package models

import (
	"time"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/google/uuid"
)

// Types
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func ResponseAPIUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdateAt:  dbUser.UpdateAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}
