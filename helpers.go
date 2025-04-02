package main

import (
	"fmt"
	"os"

	"github.com/maxBRT/gator/internal/clilogic"
	"github.com/maxBRT/gator/internal/config"
)

// checkArgs verifies that the minimum required
// command-line arguments were provided. Exits the program
// with status code 1 if insufficient arguments are supplied.
func checkArgs() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
}

// initState creates and initializes the application state
// including loading the configuration from disk.
// Returns a pointer to the initialized state object.
func initState() *clilogic.State {
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	state := &clilogic.State{
		Config: &cfg,
	}
	return state
}

// runCommandEntered parses the command-line arguments into a Command structure
// and passes it to the command handler for execution.
// Exits with status code 1 if the command execution fails.
func runCommandEntered(appState *clilogic.State, commands *clilogic.Commands) {
	// Create a Command from command-line arguments
	cmdEntered := clilogic.Command{
		Name: os.Args[1],  // First argument is the command name
		Args: os.Args[2:], // Remaining arguments are passed to the command
	}

	// Execute the command and handle any errors
	err := commands.Run(appState, cmdEntered)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
