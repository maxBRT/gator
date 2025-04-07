package config

// Package config provides configuration management for the Gator application.
// It handles reading and writing configuration data to persistent storage.

// Config represents the application's configuration structure.
// It stores database connection details and user session information.

// Config holds the application configuration settings
// that are persisted between program runs.
// It is stored as JSON in the user's home directory.
type Config struct {
	DBURL    string `json:"db_url"`            // Database connection URL
	USERNAME string `json:"current_user_name"` // Currently logged in username
}
