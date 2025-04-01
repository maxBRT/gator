package main

import (
	"fmt"

	"github.com/maxBRT/gator/internal/clilogic"
)

// main is the entry point for the Gator CLI application
// It initializes the application state, registers commands,
// and processes the command entered by the user
func main() {
	// Verify sufficient command-line arguments were provided
	checkArgs()

	// Initialize application state (config, etc.)
	appState := initState()

	// Register available commands
	commands := &clilogic.Commands{}
	commands.Register("login", clilogic.HandlerLogin)

	// Process the command entered by the user
	runCommandEntered(appState, commands)

	// Display current configuration for debugging
	fmt.Println("DB URL:", appState.Config.DBURL)
	fmt.Println("Username:", appState.Config.USERNAME)
}
