package main

import (
	"time"

	"github.com/alex-arraga/rss_project/internal/database/sqlc"
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

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	URL         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
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

// Responses
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

func resonseAPIPostForUser(dbPosts database.Post) Post {
	var description *string
	if dbPosts.Description.Valid {
		description = &dbPosts.Description.String
	}

	return Post{
		ID:          dbPosts.ID,
		CreatedAt:   dbPosts.CreatedAt,
		UpdatedAt:   dbPosts.UpdatedAt,
		Title:       dbPosts.Title,
		Description: description,
		PublishedAt: dbPosts.PublishedAt,
		URL:         dbPosts.Url,
		FeedID:      dbPosts.FeedID,
	}
}

func resonseAPIPostsForUser(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, dbPost := range dbPosts {
		posts = append(posts, resonseAPIPostForUser(dbPost))
	}

	return posts
}
