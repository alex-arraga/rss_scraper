package handlers

import (
	"net/http"

	"github.com/alex-arraga/rss_project/internal/utils"
)

func (*HandlerConfig) HealthyHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, struct{}{})
}
