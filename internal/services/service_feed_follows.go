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

func (apiCfg *ServicesConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params, err := utils.ParseRequestBody[parameters](r)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error parsing JSON: %v", err))
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
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, models.ResonseAPIFeedFollows(feedFollows))
}

func (apiCfg *ServicesConfig) HandlerGetFeedsFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollows, err := apiCfg.DB.GetFeedsFollows(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.ResonseAPIFeedsFollows(feedsFollows))
}

func (apiCfg *ServicesConfig) HandlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowId, err := utils.ParseURLParamToUUID(r, "feedFollowID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		FeedID: feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't delete feed: %v", err))
	}

	utils.RespondWithJSON(w, http.StatusOK, struct{}{})
}
