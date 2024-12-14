package api

import (
	"net/http"

	"github.com/alex-arraga/rss_project/internal/utils"
)

func (*APIConfig) HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, struct{}{})
}
