package main

import (
	"github.com/maxBRT/gator/internal/config"
)

func main() {
	// Initialize the config
	cfg := config.Config{}
	// Set the username in the config
	err := cfg.SetUser("admin")
	if err != nil {
		panic(err)
	}
	// Print the config
	println("DB URL:", cfg.DBURL)
	println("Username:", cfg.USERNAME)
}
