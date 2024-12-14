package api

import (
	"net/http"

	"github.com/alex-arraga/rss_project/internal/utils"
)

func (*APIConfig) HandlerErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusBadRequest, "Something went wrong")
}
