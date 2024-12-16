package services

import (
	"fmt"
	"net/http"
	"time"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/models"
	"github.com/alex-arraga/rss_project/internal/utils"
	"github.com/google/uuid"
)

// POST
func (apiCfg *ServicesConfig) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

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

	// Create feed in db
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.ResonseAPIFeed(feed))
}

// GET - many
func (apiCfg *ServicesConfig) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.ResonseAPIFeeds(feeds))
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
