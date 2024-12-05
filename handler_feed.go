package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/alex-arraga/rss_project/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// POST
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	params := parameters{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, resonseAPIFeed(feed))
}

// GET - many
func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	respondWithJSON(w, 200, resonseAPIFeeds(feeds))
}

// PUT
func (apiCfg *apiConfig) handlerUpdateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	params := parameter{}

	feedIDStr := chi.URLParam(r, "feedID")
	feedID, err := uuid.Parse(feedIDStr)
	if err != nil {
		log.Printf("Error parsing uuid: %v", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Couldn't read body: %v", err)
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &params)
	if err != nil {
		log.Printf("Couldn't parse body: %v", err)
	}

	feedUpdated, err := apiCfg.DB.UpdateFeed(r.Context(), database.UpdateFeedParams{
		ID:   feedID,
		Name: params.Name,
		Url:  params.URL,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't update feed: %v", err))
		return
	}

	respondWithJSON(w, 200, resonseAPIFeed(feedUpdated))
}

// DELETE - one
func (apiCfg *apiConfig) handlerDeleteFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedIDStr := chi.URLParam(r, "feedID")
	feedID, err := uuid.Parse(feedIDStr)
	if err != nil {
		log.Printf("Error parsing uuid: %v", err)
	}

	err = apiCfg.DB.DeleteFeed(r.Context(), feedID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})
}
