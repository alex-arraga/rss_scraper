package handlers

import (
	"fmt"
	"net/http"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/utils"
	"github.com/google/uuid"
)

func (h *HandlerConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params, err := utils.ParseRequestBody[parameters](r)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollows, err := h.Services.CreateFeedFollow(r.Context(), user.ID, params.FeedID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, feedFollows)
}

func (h *HandlerConfig) HandlerGetFeedsFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollows, err := h.Services.GetFeedsFollows(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, feedsFollows)
}

func (apiCfg *APIConfig) HandlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
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
