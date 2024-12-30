package handlers

import (
	"fmt"
	"net/http"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/utils"
)

// POST
func (h *HandlerConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
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
	feed, err := h.Container.FeedService.CreateFeed(r.Context(), user.ID, params.Name, params.URL)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, feed)
}

// GET - many
func (h *HandlerConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := h.Container.FeedService.GetFeeds(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, feeds)
}

// PUT
func (h *HandlerConfig) HandlerUpdateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
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
	feedUpdated, err := h.Container.FeedService.UpdateFeed(r.Context(), feedID, params.Name, params.URL)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't update feed: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, feedUpdated)
}

// DELETE - one
func (h *HandlerConfig) HandlerDeleteFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedID, err := utils.ParseURLParamToUUID(r, "feedID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	err = h.Container.FeedService.DeleteFeed(r.Context(), feedID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, struct{}{})
}
