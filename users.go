package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/crayboi420/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type userStruct struct {
		Name string `json:"name"`
	}
	usr := userStruct{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&usr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
	}

	name := usr.Name
	uuid := uuid.New()
	user := database.CreateUserParams{
		Name:      name,
		ID:        uuid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	returned, err := cfg.DB.CreateUser(r.Context(), user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't add to database : "+err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, returned)
}

func (cfg *apiConfig) handlerGetUserApi(w http.ResponseWriter, r *http.Request) {
	apiKey := strings.TrimLeft(r.Header.Get("Authorization"), "ApiKey ")

	retr, err := cfg.DB.GetUserApi(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find ApiKey")
		return
	}
	respondWithJSON(w, http.StatusOK, retr)
}
