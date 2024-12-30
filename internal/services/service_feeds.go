package services

import (
	"context"
	"fmt"
	"time"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/models"
	"github.com/google/uuid"
)

type FeedDatabase interface {
	CreateFeed(ctx context.Context, params database.CreateFeedParams) (database.Feed, error)
	GetFeeds(ctx context.Context) ([]database.Feed, error)
	UpdateFeed(ctx context.Context, params database.UpdateFeedParams) (database.Feed, error)
	DeleteFeed(ctx context.Context, feedID uuid.UUID) error
}

type FeedService struct {
	DB FeedDatabase
}

// POST
func (fs *FeedService) CreateFeed(ctx context.Context, userID uuid.UUID, name string, url string) (models.Feed, error) {
	// Create feed in db
	feed, err := fs.DB.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    userID,
	})
	if err != nil {
		return models.Feed{}, fmt.Errorf("couldn't create feed: %v", err)
	}

	return models.ResonseAPIFeed(feed), nil
}

// GET - many
func (fs *FeedService) GetFeeds(ctx context.Context) ([]models.Feed, error) {
	feeds, err := fs.DB.GetFeeds(ctx)
	if err != nil {
		return []models.Feed{}, fmt.Errorf("couldn't get feeds: %v", err)
	}

	return models.ResonseAPIFeeds(feeds), nil
}

// PUT
func (fs *FeedService) UpdateFeed(ctx context.Context, feedID uuid.UUID, name string, url string) (models.Feed, error) {
	feedUpdated, err := fs.DB.UpdateFeed(ctx, database.UpdateFeedParams{
		ID:   feedID,
		Name: name,
		Url:  url,
	})
	if err != nil {
		return models.Feed{}, fmt.Errorf("couldn't update feed: %v", err)
	}

	return models.ResonseAPIFeed(feedUpdated), nil
}

// DELETE - one
func (fs *FeedService) DeleteFeed(ctx context.Context, feedID uuid.UUID) error {
	err := fs.DB.DeleteFeed(ctx, feedID)
	if err != nil {
		return fmt.Errorf("couldn't delete feed: %v", err)
	}

	return nil
}
