package clilogic

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/maxBRT/gator/internal/database"
	"github.com/maxBRT/gator/internal/rss"
)

// scrapeFeeds handles the feed scraping process:
// 1. Gets the next feed that needs updating
// 2. Marks it as fetched
// 3. Retrieves the feed content
// 4. Processes and stores new entries
//
// The function implements an idempotent update process that:
// - Skips duplicate entries
// - Handles parsing errors gracefully
// - Updates feed fetch timestamps
func scrapeFeeds(state *State) error {
	feed := database.Feed{}
	rssFeed := &rss.RSSFeed{}

	feed, err := state.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = state.DB.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feed.ID,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	rssFeed, err = rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		pubTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			// Use current time as fallback if parsing fails
			pubTime = time.Now()
		}
		_, err = state.DB.AddPost(context.Background(), database.AddPostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubTime,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}

	}

	return nil
}
