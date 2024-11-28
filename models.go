package main

import (
	"time"

	"github.com/alex-arraga/rss_project/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	User_ID   uuid.UUID `json:"user_id"`
}
type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Feed_ID   uuid.UUID `json:"feed_id"`
	User_ID   uuid.UUID `json:"user_id"`
}

func responseAPIUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdateAt:  dbUser.UpdateAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

func resonseAPIFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdateAt:  dbFeed.UpdateAt,
		Name:      dbFeed.Name,
		URL:       dbFeed.Url,
		User_ID:   dbFeed.UserID,
	}
}

func resonseAPIFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, resonseAPIFeed(dbFeed))
	}
	return feeds
}

func resonseAPIFeedFollows(dbFeedFollows database.FeedFollow) FeedFollows {
	return FeedFollows{
		ID:        dbFeedFollows.ID,
		CreatedAt: dbFeedFollows.CreatedAt,
		UpdateAt:  dbFeedFollows.UpdateAt,
		Feed_ID:   dbFeedFollows.FeedID,
		User_ID:   dbFeedFollows.UserID,
	}
}

func resonseAPIFeedsFollows(dbFeedsFollows []database.FeedFollow) []FeedFollows {
	feedFollows := []FeedFollows{}
	for _, dbFeedFollow := range dbFeedsFollows {
		feedFollows = append(feedFollows, resonseAPIFeedFollows(dbFeedFollow))
	}

	return feedFollows
}
