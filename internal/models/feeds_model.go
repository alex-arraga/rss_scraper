package models

import (
	"time"

	"github.com/google/uuid"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	User_ID   uuid.UUID `json:"user_id"`
}

func ResonseAPIFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdateAt:  dbFeed.UpdateAt,
		Name:      dbFeed.Name,
		URL:       dbFeed.Url,
		User_ID:   dbFeed.UserID,
	}
}

// Responses
func ResonseAPIFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, ResonseAPIFeed(dbFeed))
	}
	return feeds
}
