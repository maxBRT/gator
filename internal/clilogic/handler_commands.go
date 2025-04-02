package clilogic

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/maxBRT/gator/internal/database"
)

// HandlerRegister - Where new Gators come to get their student ID
func HandlerRegister(State *State, cmd Command) error {
	// Ensure username was provided
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is empty")
	}

	// Create a fancy new user with all the bells and whistles
	params := database.CreateUserParams{
		ID:        uuid.New(), // A unique ID as special as a snowflake
		CreatedAt: time.Now(), // Born at this exact moment - historic!
		UpdatedAt: time.Now(), // Fresh out of the digital egg
		Name:      cmd.Args[0],
	}

	// Check if this Gator name is already taken in the swamp
	_, err := State.DB.GetUser(context.Background(), params.Name)
	if err == nil {
		fmt.Println("User already exist")
		os.Exit(1)
	}

	// Add the new Gator to our carefully curated collection
	usr, err := State.DB.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Update the config so we know who's the boss Gator now
	err = State.Config.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to set username: %w", err)
	}

	// Notify the user of successful login
	// Pop the champagne! We have a new Gator in town!
	fmt.Printf("User: %s succesfully added to the database \n", cmd.Args[0])
	fmt.Println(usr)
	return nil
}

// HandlerLogin - Signs in an existing Gator by checking if the username exists
// and setting it as the current user in the application configuration
func HandlerLogin(State *State, cmd Command) error {
	// Ensure username was provided
	// Can't login as nobody - we're not a ghost app!
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is empty")
	}

	// Check if this Gator actually exists in our records
	_, err := State.DB.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Println("User does not exist") // Who are you? You're not in the Gator registry!
		os.Exit(1)
	}

	// Update the configuration with the provided username
	// Crown the new king/queen of the swamp!
	err = State.Config.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to set username: %w", err) // Failed to update the Gator throne
	}

	// Notify the user of successful login
	// The Gator has been recognized by the council!
	fmt.Printf("Username set to: %s\n", cmd.Args[0])
	return nil
}

// HandlerReset - The nuclear option for when things go terribly wrong
// "Drain the swamp! Start fresh! It's not a bug, it's a feature!"
func HandlerReset(state *State, cmd Command) error {
	err := state.DB.ResetUsersTable(context.Background())
	if err != nil {
		return fmt.Errorf("unable to truncate users table %w", err) // Even our reset button is broken!
	}
	fmt.Println("Users table succesfully reseted.") // Goodbye Gators, hello empty swamp!
	return nil
}

// HandlerUsers - The Gator census bureau
// "Let's count all the Gators in the swamp!"
func HandlerUsers(state *State, cmd Command) error {
	users, err := state.DB.GetUsers(context.Background())
	if err != nil {
		return err // Failed to count Gators - they're too slippery!
	}
	for _, user := range users {
		if user == state.Config.USERNAME {
			fmt.Printf("* %s (current)\n", user) // This Gator is YOU! (Look in the mirror)
		} else {
			fmt.Printf("* %s\n", user) // Just another Gator in the swamp
		}
	}
	return nil
}
