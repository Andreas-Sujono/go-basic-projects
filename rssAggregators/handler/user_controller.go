package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/andreas-sujono/go-basic-projects/rssAggregators/internal/database"
	"github.com/google/uuid"
)

type ApiConfig struct {
	DB *database.Queries
}

func (config *ApiConfig) GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, ToUser(user))
}

func (config *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid parameters")
		return
	}

	userData := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	}

	user, err := config.DB.CreateUser(r.Context(), userData)

	if err != nil {
		log.Println("Error creating user")
		respondWithError(w, http.StatusInternalServerError, "Error creating user")
	}

	respondWithJSON(w, http.StatusOK,
		ToUser(user),
	)
}
