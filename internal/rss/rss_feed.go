package rss

// Package rss provides structures and functions for handling RSS feed data.
// It defines the core data types used to represent RSS feeds and their entries.

// RSSFeed represents an RSS feed's structure according to the RSS 2.0 specification.
// It contains channel information and a collection of feed items.
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

// RSSItem represents a single entry in an RSS feed containing
// essential information about the content
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
