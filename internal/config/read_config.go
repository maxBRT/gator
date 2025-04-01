package config

import (
	"encoding/json"
	"os"
)

func ReadConfig() (Config, error) {
	path := setFilePath()
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		file, err = os.Create(path)
		if err != nil {
			return Config{}, err
		}
	}
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	config := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}
