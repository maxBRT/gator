package clilogic

// Package clilogic implements the command-line interface logic for the Gator application.
// It provides command handling, state management, and business logic implementation.

import (
	"github.com/maxBRT/gator/internal/config"
	"github.com/maxBRT/gator/internal/database"
)

// State maintains the runtime state of the application including:
// - Database connection and queries
// - Configuration settings
// - Current user session information

// State holds the application's runtime state
// It maintains references to configuration and other
// stateful components needed during program execution
type State struct {
	DB     *database.Queries
	Config *config.Config // Reference to the application configuration
}
