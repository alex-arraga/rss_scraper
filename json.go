package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error: ", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func parseRequestBody[T any](r *http.Request) (T, error) {
	var params T

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&params); err != nil {
		return params, err
	}
	return params, nil
}

func parseURLParamToUUID(r *http.Request, urlParam string) (uuid.UUID, error) {
	reqID := chi.URLParam(r, urlParam)
	parsedID, err := uuid.Parse(reqID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID: %v", err)
	}

	return parsedID, nil
}
