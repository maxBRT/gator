package clilogic

import (
	"github.com/maxBRT/gator/internal/config"
	"github.com/maxBRT/gator/internal/database"
)

// State holds the application's runtime state
// It maintains references to configuration and other
// stateful components needed during program execution
type State struct {
	DB     *database.Queries
	Config *config.Config // Reference to the application configuration
}
