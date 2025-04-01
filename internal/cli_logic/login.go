package clilogic

import (
	"fmt"
)

func handlerLogin(state *State, cmd command) error {
	// Check if the username is empty
	if len(cmd.args) == 0 {
		return fmt.Errorf("username is empty")
	}
	// Set the username in the config
	err := state.config.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to set username: %w", err)
	}
	// Print the username
	fmt.Printf("Username set to: %s\n", cmd.args[0])
	return nil
}
