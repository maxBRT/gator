package config

// Config represents the application configuration stored in JSON format
type Config struct {
	DBURL    string `json:"db_url"`            // Database connection URL
	USERNAME string `json:"current_user_name"` // Currently logged in username
}
