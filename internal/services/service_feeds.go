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
func (apiCfg *ServicesConfig) UpdateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	// Get and validate feed id
	feedID, err := utils.ParseURLParamToUUID(r, "feedID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Parse to JSON
	params, err := utils.ParseRequestBody[parameters](r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid input: %v", err))
		return
	}

	// Validate params
	if params.Name == "" || params.URL == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Name and URL are required")
		return
	}

	// Update feed in db
	feedUpdated, err := apiCfg.DB.UpdateFeed(r.Context(), database.UpdateFeedParams{
		ID:   feedID,
		Name: params.Name,
		Url:  params.URL,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't update feed: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.ResonseAPIFeed(feedUpdated))
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
