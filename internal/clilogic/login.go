package clilogic

import (
	"context"
	"fmt"
	"os"
)

// HandlerLogin processes the login command
// It validates the username argument and updates the configuration
// Returns an error if the username is empty or if setting the username fails
func HandlerLogin(State *State, cmd Command) error {
	// Ensure username was provided
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is empty")
	}
	_, err := State.DB.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Println("User does not exist")
		os.Exit(1)
	}

	// Update the configuration with the provided username
	err = State.Config.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to set username: %w", err)
	}

	// Notify the user of successful login
	fmt.Printf("Username set to: %s\n", cmd.Args[0])
	return nil
}
