package rss

// Package rss provides structures and functions for handling RSS feed data.
// It defines the core data types used to represent RSS feeds and their entries.

// RSSFeed represents an RSS feed's structure according to the RSS 2.0 specification.
// It contains channel information and a collection of feed items.
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`       // Feed title from the channel
		Link        string    `xml:"link"`        // Website URL associated with the feed
		Description string    `xml:"description"` // Feed description/summary
		Item        []RSSItem `xml:"item"`        // Collection of feed entries
	} `xml:"channel"`
}

// RSSItem represents a single entry in an RSS feed containing
// essential information about the content. Each item typically
// represents a blog post, news article, or other content unit.
type RSSItem struct {
	Title       string `xml:"title"`       // Title of the content entry
	Link        string `xml:"link"`        // URL to the full content
	Description string `xml:"description"` // Summary or full content of the entry
	PubDate     string `xml:"pubDate"`     // Publication date in RFC1123Z format
}
