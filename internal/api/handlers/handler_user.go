package handlers

import (
	"fmt"
	"net/http"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/models"
	"github.com/alex-arraga/rss_project/internal/utils"
)

func (h *HandlerConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name   string `json:"name"`
		ApiKey string `json:"api_key"`
	}

	params, err := utils.ParseRequestBody[parameters](r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid input: %v", err))
		return
	}

	if params.Name == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Name is required")
		return
	}

	user, err := h.Container.UserService.CreateUser(r.Context(), params.Name)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}

func (h *HandlerConfig) HandlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseAPIUser(user))
}

func (h *HandlerConfig) HandlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := h.Container.UserService.GetPostsForUser(r.Context(), user.ID, 10)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't get the posts %v ", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, posts)
}
