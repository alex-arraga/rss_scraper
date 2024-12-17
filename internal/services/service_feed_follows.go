package services

import (
	"context"
	"fmt"
	"time"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/models"
	"github.com/google/uuid"
)

func (srv *ServicesConfig) CreateFeedFollow(ctx context.Context, userID, feedID uuid.UUID) (models.FeedFollows, error) {
	feedFollows, err := srv.DB.CreateFeedFollows(ctx, database.CreateFeedFollowsParams{
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

func (srv *ServicesConfig) GetFeedsFollows(ctx context.Context, userID uuid.UUID) ([]models.FeedFollows, error) {
	feedsFollows, err := srv.DB.GetFeedsFollows(ctx, userID)
	if err != nil {
		return []models.FeedFollows{}, fmt.Errorf("couldn't get feeds: %v", err)
	}

	return models.ResonseAPIFeedsFollows(feedsFollows), nil
}

func (srv *ServicesConfig) DeleteFeedFollows(ctx context.Context, userID, feedID uuid.UUID) error {
	err := srv.DB.DeleteFeedFollows(ctx, database.DeleteFeedFollowsParams{
		FeedID: feedID,
		UserID: userID,
	})
	if err != nil {
		return fmt.Errorf("couldn't delete feed: %v", err)
	}

	return nil
}
