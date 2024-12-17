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
