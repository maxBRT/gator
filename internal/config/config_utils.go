package config

import (
	"encoding/json"
	"os"
)

const configFileName = "/.gatorconfig.json"
const db_url = "postgres://example"

func setFilePath() string {
	home_path, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return home_path + configFileName
}

func writeConfig(c *Config) {
	path := setFilePath()
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(c)
	if err != nil {
		return
	}
}
