package rss

// Package rss provides functionality for fetching and parsing RSS feeds.
// It handles HTTP requests, XML parsing, and content sanitization.

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

// FetchFeed retrieves and parses an RSS feed from a provided URL.
// Parameters:
//   - ctx: Context for request cancellation and timeouts
//   - feedURL: The URL of the RSS feed to fetch
//
// Returns:
//   - *RSSFeed: Parsed feed data structure
//   - error: Any error encountered during fetching or parsing
func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	rssFeed := &RSSFeed{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return rssFeed, err
	}
	req.Header.Set("User-Agent", "gator")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return rssFeed, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return rssFeed, err
	}
	xml.Unmarshal(body, rssFeed)
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for i, item := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Description = html.UnescapeString(item.Description)
		rssFeed.Channel.Item[i].Title = html.UnescapeString(item.Title)
	}

	return rssFeed, nil
}
