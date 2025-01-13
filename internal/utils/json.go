package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error: ", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing data: %v", err)
	}
}

func ParseRequestBody[T any](r *http.Request) (T, error) {
	var params T

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&params); err != nil {
		return params, err
	}
	return params, nil
}

// Convert string URL param to UUID
func ParseURLParamToUUID(r *http.Request, urlParam string) (uuid.UUID, error) {
	reqID := chi.URLParam(r, urlParam)
	parsedID, err := uuid.Parse(reqID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID: %v", err)
	}

	return parsedID, nil
}
