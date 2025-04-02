package config

import (
	"encoding/json"
	"os"
)

// ReadConfig loads the application configuration from the config file
// If the file doesn't exist, it creates an empty one
func ReadConfig() (Config, error) {
	path := setFilePath()
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		// Create the config file if it doesn't exist
		file, err = os.Create(path)
		if err != nil {
			return Config{}, err // Failed to create a new diary - the Gator is sad
		}
	}
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	// Parse the JSON into a Config struct
	// Translating from Gator-JSON to Gator-Go
	config := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}
