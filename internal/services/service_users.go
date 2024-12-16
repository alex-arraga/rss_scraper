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

func (userSrv *ServicesConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	user, err := userSrv.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.ResponseAPIUser(user))
}

func (userSrv *ServicesConfig) GetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseAPIUser(user))
}

func (userSrv *ServicesConfig) GetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := userSrv.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't get the posts %v ", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.ResonseAPIPostsForUser(posts))
}
