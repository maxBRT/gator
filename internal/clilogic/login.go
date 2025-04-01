package clilogic

import (
	"fmt"
)

func HandlerLogin(State *State, cmd Command) error {
	// Check if the username is empty
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is empty")
	}
	// Set the username in the config
	err := State.Config.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to set username: %w", err)
	}
	// Print the username
	fmt.Printf("Username set to: %s\n", cmd.Args[0])
	return nil
}
