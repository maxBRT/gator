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
//
// State is passed to command handlers to provide access to these
// shared resources in a structured way.
type State struct {
	DB     *database.Queries // Database query interface
	Config *config.Config    // Application configuration
}
