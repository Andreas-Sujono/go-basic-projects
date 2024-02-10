package handler

import (
	"net/http"
	"strconv"

	"github.com/andreas-sujono/go-basic-projects/rssAggregators/internal/database"
)

func (config *ApiConfig) GetPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid limit")
		return
	}

	posts, err := config.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Cannot get posts for user")
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
