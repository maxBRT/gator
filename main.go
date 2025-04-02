package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/maxBRT/gator/internal/clilogic"
	"github.com/maxBRT/gator/internal/config"
	"github.com/maxBRT/gator/internal/database"
)

// main is the entry point for the Gator CLI application
// It initializes the application state, registers commands,
// and processes the command entered by the user
func main() {
	// Verify sufficient command-line arguments were provided
	checkArgs()

	// Initialize application state (config, etc.)
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	appState := &clilogic.State{
		Config: &cfg,
	}

	db, err := sql.Open("postgres", appState.Config.DBURL)
	if err != nil {
		fmt.Println(err)

	}
	dbQueries := database.New(db)
	appState.DB = dbQueries

	// Register available commands
	commands := &clilogic.Commands{}
	commands.Register("login", clilogic.HandlerLogin)
	commands.Register("register", clilogic.HandlerRegister)

	// Process the command entered by the user
	runCommandEntered(appState, commands)

	// Display current configuration for debugging
	fmt.Println("DB URL:", appState.Config.DBURL)
}
