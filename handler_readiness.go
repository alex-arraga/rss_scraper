package main

import (
	"net/http"

	"github.com/alex-arraga/rss_project/internal/utils"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, struct{}{})
}
