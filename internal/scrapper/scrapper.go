package scrapper

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/google/uuid"
)

func StartScrapping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scrapping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)

	// Immediate execution loop
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds:", err)
			continue
		}

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

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched:", err)
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed:", err)
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		var pubAt time.Time
		if item.PubDate == "" {
			pubAt = time.Now().UTC()
		} else {
			pubAt, err = time.Parse(time.RFC1123Z, item.PubDate)
			if err != nil {
				log.Printf("couldn't parse date: %v - err: %v", item.PubDate, err)
				continue
			}
		}

		_, err = db.CreatePost(context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Description: description,
				Url:         item.Link,
				FeedID:      feed.ID,
				PublishedAt: pubAt,
			})
		if err != nil {
			possibleErrs := []string{"duplicate key", "llave duplicada"}
			skipError := false

			for _, possibleErr := range possibleErrs {
				if strings.Contains(err.Error(), possibleErr) {
					skipError = true
					break
				}
			}

			if skipError {
				continue
			}

			log.Printf("failed creating new post: %v", err)
		}
	}

	log.Printf("Feed %s collected, %v post found", feed.Name, len(rssFeed.Channel.Item))
}
