package handler

import (
	"net/http"

	"github.com/andreas-sujono/go-basic-projects/rssAggregators/internal/auth"
	"github.com/andreas-sujono/go-basic-projects/rssAggregators/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (config *ApiConfig) AuthMiddleware(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Api key not found")
			return
		}

		user, err := config.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		handler(w, r, user)
	}
}
