package clilogic

import (
	"github.com/maxBRT/gator/internal/config"
)

// State holds the application's runtime state
// It maintains references to configuration and other
// stateful components needed during program execution
type State struct {
	Config *config.Config // Reference to the application configuration
}
