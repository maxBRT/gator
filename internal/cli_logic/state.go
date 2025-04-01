package clilogic

import (
	"github.com/maxBRT/gator/internal/config"
)

type State struct {
	// The current state of the application
	config *config.Config
}
