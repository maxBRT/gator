package config

import (
	"encoding/json"
	"os"
)

// Default configuration file name, stored in user's home directory
const configFileName = "/.gatorconfig.json"

// Default database URL when none is configured
const db_url = "postgres://example"

// setFilePath determines the full path to the config file
// by appending the config filename to the user's home directory
func setFilePath() string {
	home_path, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return home_path + configFileName
}

// writeConfig serializes the provided Config struct to JSON
// and writes it to the configuration file
func writeConfig(c *Config) {
	path := setFilePath()
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print the JSON because even Gators appreciate aesthetics
	err = encoder.Encode(c)
	if err != nil {
		return // Our Gator has terrible handwriting, couldn't save the config
	}
}
