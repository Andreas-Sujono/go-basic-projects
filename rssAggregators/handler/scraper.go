package handler

import (
	"context"
	"database/sql"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/andreas-sujono/go-basic-projects/rssAggregators/internal/database"
	"github.com/google/uuid"
)

func StartScraping(
	db *database.Queries,
	concurrency int32,
	duration time.Duration,
) {

	ticker := time.NewTicker(duration)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), concurrency)
		if err != nil {
			log.Println("Failed get next feeds")
			continue
		}

		log.Printf("Found next %v feeds to fetch", len(feeds))
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	db.MarkFeedFetched(context.Background(), feed.ID)

	feedData, err := fetchFeed(feed.Url)
	if err != nil {
		log.Println("Failed to fetch feed: ", err)
		return
	}

	for _, feedItem := range feedData.Channel.Item {
		pubDateParsed, pubDateParsedErr := time.Parse(time.RFC1123, feedItem.PubDate)
		_, err := db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID:        uuid.New(),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
				Title:     feedItem.Title,
				Url:       feedItem.Link,
				Description: sql.NullString{
					String: feedItem.Description,
					Valid:  true,
				},
				PublishedAt: sql.NullTime{
					Time:  pubDateParsed,
					Valid: pubDateParsedErr == nil,
				},
				FeedID: feed.ID,
			},
		)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}

}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(feedUrl string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	res, err := httpClient.Get(feedUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	return &feed, nil
}
