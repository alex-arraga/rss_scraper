package models

import (
	"time"

	"github.com/google/uuid"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
)

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

func ResonseAPIPostForUser(dbPosts database.Post) Post {
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

func ResonseAPIPostsForUser(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, dbPost := range dbPosts {
		posts = append(posts, ResonseAPIPostForUser(dbPost))
	}

	return posts
}
