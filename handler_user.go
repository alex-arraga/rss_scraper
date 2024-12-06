package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name   string `json:"name"`
		ApiKey string `json:"api_key"`
	}

	params, err := parseRequestBody[parameters](r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid input: %v", err))
		return
	}

	if params.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Name is required")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, responseAPIUser(user))
}

func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, responseAPIUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't get the posts %v ", err))
		return
	}

	respondWithJSON(w, http.StatusOK, resonseAPIPostsForUser(posts))
}
