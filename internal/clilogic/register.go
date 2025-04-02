package clilogic

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/maxBRT/gator/internal/database"
)

func HandlerRegister(State *State, cmd Command) error {
	// Ensure username was provided
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is empty")
	}

	// Update the configuration with the provided username
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	_, err := State.DB.GetUser(context.Background(), params.Name)
	if err == nil {
		fmt.Println("User already exist")
		os.Exit(1)
	}

	usr, err := State.DB.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Notify the user of successful login
	fmt.Printf("User: %s succesfully added to the database \n", cmd.Args[0])
	fmt.Println(usr)
	HandlerLogin(State, cmd)
	return nil
}
