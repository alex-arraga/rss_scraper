package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/google/uuid"
)

// POST
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	params, err := parseRequestBody[parameters](r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid input: %v", err))
		return
	}

	// Validate params
	if params.Name == "" || params.URL == "" {
		respondWithError(w, http.StatusBadRequest, "Name and URL are required")
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
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, resonseAPIFeed(feed))
}

// GET - many
func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, resonseAPIFeeds(feeds))
}

// PUT
func (apiCfg *apiConfig) handlerUpdateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	// Get and validate feed id
	feedID, err := parseURLParamToUUID(r, "feedID")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Parse to JSON
	params, err := parseRequestBody[parameters](r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid input: %v", err))
		return
	}

	// Validate params
	if params.Name == "" || params.URL == "" {
		respondWithError(w, http.StatusBadRequest, "Name and URL are required")
		return
	}

	// Update feed in db
	feedUpdated, err := apiCfg.DB.UpdateFeed(r.Context(), database.UpdateFeedParams{
		ID:   feedID,
		Name: params.Name,
		Url:  params.URL,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't update feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, resonseAPIFeed(feedUpdated))
}

// DELETE - one
func (apiCfg *apiConfig) handlerDeleteFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedID, err := parseURLParamToUUID(r, "feedID")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	err = apiCfg.DB.DeleteFeed(r.Context(), feedID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't delete feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
