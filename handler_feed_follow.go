package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params, err := parseRequestBody[parameters](r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollows, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, resonseAPIFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerGetFeedsFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollows, err := apiCfg.DB.GetFeedsFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, resonseAPIFeedsFollows(feedsFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowId, err := parseURLParamToUUID(r, "feedFollowID")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		FeedID: feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't delete feed: %v", err))
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
