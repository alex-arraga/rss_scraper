package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/models"
)

type FeedFollowsDatabase interface {
	CreateFeedFollows(ctx context.Context, params database.CreateFeedFollowsParams) (database.FeedFollow, error)
	GetFeedsFollows(ctx context.Context, userID uuid.UUID) ([]database.FeedFollow, error)
	DeleteFeedFollows(ctx context.Context, params database.DeleteFeedFollowsParams) error
}

type FeedFollowService struct {
	DB FeedFollowsDatabase
}

func (fs *FeedFollowService) CreateFeedFollow(ctx context.Context, userID, feedID uuid.UUID) (models.FeedFollows, error) {
	feedFollows, err := fs.DB.CreateFeedFollows(ctx, database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		FeedID:    feedID,
		UserID:    userID,
	})
	if err != nil {
		return models.FeedFollows{}, fmt.Errorf("couldn't create feed: %v", err)
	}

	return models.ResonseAPIFeedFollows(feedFollows), nil
}

func (fs *FeedFollowService) GetFeedsFollows(ctx context.Context, userID uuid.UUID) ([]models.FeedFollows, error) {
	feedsFollows, err := fs.DB.GetFeedsFollows(ctx, userID)
	if err != nil {
		return []models.FeedFollows{}, fmt.Errorf("couldn't get feeds: %v", err)
	}

	return models.ResonseAPIFeedsFollows(feedsFollows), nil
}

func (fs *FeedFollowService) DeleteFeedFollows(ctx context.Context, userID, feedID uuid.UUID) error {
	err := fs.DB.DeleteFeedFollows(ctx, database.DeleteFeedFollowsParams{
		FeedID: feedID,
		UserID: userID,
	})
	if err != nil {
		return fmt.Errorf("couldn't delete feed: %v", err)
	}

	return nil
}
