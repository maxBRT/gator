package config

// SetUser updates the username in the configuration
// It reads the current config, updates the username,
// and writes the updated config back to the file
func (c *Config) SetUser(username string) error {
	// Read current configuration
	configData, err := ReadConfig()
	if err != nil {
		return err
	}

	// Ensure the database URL is set
	if configData.DBURL == "" {
		configData.DBURL = db_url
	} else {
		c.DBURL = configData.DBURL
	}

	// Update username in the config
	c.USERNAME = username

	// Save updated configuration
	writeConfig(c)
	return nil
}
