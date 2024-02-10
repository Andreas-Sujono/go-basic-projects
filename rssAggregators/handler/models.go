package handler

import (
	"database/sql"
	"time"

	"github.com/andreas-sujono/go-basic-projects/rssAggregators/internal/database"
	"github.com/google/uuid"
)

type User struct {
	Name      string    `json:"name"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func ToUser(dbUser database.User) User {
	return User{
		Name:      dbUser.Name,
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Url           string     `json:"url"`
	UserId        uuid.UUID  `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func ToFeed(dbFeed database.Feed) Feed {
	return Feed{
		Name:          dbFeed.Name,
		ID:            dbFeed.ID,
		CreatedAt:     dbFeed.CreatedAt,
		UpdatedAt:     dbFeed.UpdatedAt,
		Url:           dbFeed.Url,
		UserId:        dbFeed.UserID,
		LastFetchedAt: sqlNullTimeToTime(dbFeed.LastFetchedAt),
	}
}

func ToFeeds(dbFeeds []database.Feed) []Feed {
	res := make([]Feed, len(dbFeeds))
	for i, feed := range dbFeeds {
		res[i] = ToFeed(feed)
	}

	return res
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func ToFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
	}
}

func ToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	result := make([]FeedFollow, len(feedFollows))
	for i, feedFollow := range feedFollows {
		result[i] = ToFeedFollow(feedFollow)
	}
	return result
}

type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func ToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: sqlNullStringToStringPtr(post.Description),
		PublishedAt: sqlNullTimeToTime(post.PublishedAt),
		FeedID:      post.FeedID,
	}
}

func ToPosts(posts []database.Post) []Post {
	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = ToPost(post)
	}
	return result
}

func sqlNullTimeToTime(time sql.NullTime) *time.Time {
	if time.Valid {
		return &time.Time
	}

	return nil
}

func sqlNullStringToStringPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}
