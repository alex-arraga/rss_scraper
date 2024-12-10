package main

import (
	"net/http"

	"github.com/alex-arraga/rss_project/internal/utils"
)

func handlerErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusBadRequest, "Something went wrong")
}
