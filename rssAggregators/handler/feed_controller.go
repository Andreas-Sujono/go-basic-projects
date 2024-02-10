package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/andreas-sujono/go-basic-projects/rssAggregators/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      params.Name,
		Url:       params.URL,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	log.Printf("feed %v,\n feedFoolow %v,\n combined %v", ToFeed(feed), ToFeedFollow(feedFollow), struct {
		feedFollow FeedFollow
	}{
		feedFollow: ToFeedFollow(feedFollow),
	})
	respondWithJSON(w, http.StatusOK, struct {
		Feed       Feed
		FeedFollow FeedFollow
	}{
		Feed:       ToFeed(feed),
		FeedFollow: ToFeedFollow(feedFollow),
	})
}

func (cfg *ApiConfig) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	respondWithJSON(w, http.StatusOK, ToFeeds(feeds))
}
