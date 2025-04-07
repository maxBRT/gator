package clilogic

import (
	"context"

	"github.com/maxBRT/gator/internal/database"
)

// MiddlewareLoggedIn wraps command handlers to ensure user authentication.
// It verifies the current user exists in the database before executing
// protected commands.
//
// Parameters:
//   - handler: The protected command handler function
//
// Returns:
//   - A wrapped handler that includes authentication checks
func MiddlewareLoggedIn(handler func(state *State, cmd Command, user database.User) error) func(state *State, cmd Command) error {
	return func(state *State, cmd Command) error {
		// Example implementation: Retrieve the user and call the handler
		user, err := state.DB.GetUser(context.Background(), state.Config.USERNAME)
		if err != nil {
			return err
		}
		return handler(state, cmd, user)
	}
}
