package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/models"
	"github.com/alex-arraga/rss_project/internal/utils"
	"github.com/google/uuid"
)

// POST
func (srv *ServicesConfig) CreateFeed(ctx context.Context, userID uuid.UUID, name string, url string) (models.Feed, error) {
	// Create feed in db
	feed, err := srv.DB.CreateFeed(ctx, database.CreateFeedParams{
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
func (srv *ServicesConfig) GetFeeds(ctx context.Context) ([]models.Feed, error) {
	feeds, err := srv.DB.GetFeeds(ctx)
	if err != nil {
		return []models.Feed{}, fmt.Errorf("couldn't get feeds: %v", err)
	}

	return models.ResonseAPIFeeds(feeds), nil
}

// PUT
func (srv *ServicesConfig) UpdateFeed(ctx context.Context, feedID uuid.UUID, name string, url string) (models.Feed, error) {
	feedUpdated, err := srv.DB.UpdateFeed(ctx, database.UpdateFeedParams{
		ID:   feedID,
		Name: name,
		Url:  url,
	})
	if err != nil {
		return models.Feed{}, fmt.Errorf("Couldn't update feed: %v", err)
	}

	return models.ResonseAPIFeed(feedUpdated), nil
}

// DELETE - one
func (apiCfg *ServicesConfig) DeleteFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedID, err := utils.ParseURLParamToUUID(r, "feedID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	err = apiCfg.DB.DeleteFeed(r.Context(), feedID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't delete feed: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, struct{}{})
}
