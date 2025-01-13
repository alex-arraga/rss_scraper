package models

import (
	"time"

	"github.com/google/uuid"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
)

type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Feed_ID   uuid.UUID `json:"feed_id"`
	User_ID   uuid.UUID `json:"user_id"`
}

func ResonseAPIFeedFollows(dbFeedFollows database.FeedFollow) FeedFollows {
	return FeedFollows{
		ID:        dbFeedFollows.ID,
		CreatedAt: dbFeedFollows.CreatedAt,
		UpdateAt:  dbFeedFollows.UpdateAt,
		Feed_ID:   dbFeedFollows.FeedID,
		User_ID:   dbFeedFollows.UserID,
	}
}

func ResonseAPIFeedsFollows(dbFeedsFollows []database.FeedFollow) []FeedFollows {
	feedFollows := []FeedFollows{}
	for _, dbFeedFollow := range dbFeedsFollows {
		feedFollows = append(feedFollows, ResonseAPIFeedFollows(dbFeedFollow))
	}

	return feedFollows
}
