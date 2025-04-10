package main

// Package main is the entry point for the Gator RSS feed aggregator CLI application.
// It provides functionality for user management, feed subscriptions, and content aggregation.

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/maxBRT/gator/internal/clilogic"
	"github.com/maxBRT/gator/internal/config"
	"github.com/maxBRT/gator/internal/database"
)

// main is the entry point for the Gator CLI application.
// It initializes the application state, registers commands,
// and processes the command entered by the user.
func main() {
	// Verify sufficient command-line arguments were provided
	checkArgs()

	// Initialize application state from configuration
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	appState := &clilogic.State{
		Config: &cfg,
	}

	// Establish database connection
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
	commands.Register("reset", clilogic.HandlerReset)
	commands.Register("users", clilogic.HandlerUsers)
	commands.Register("agg", clilogic.HandlerAggregate)
	commands.Register("addfeed", clilogic.MiddlewareLoggedIn(clilogic.HandlerAddFeed))
	commands.Register("feeds", clilogic.HandlerGetFeeds)
	commands.Register("follow", clilogic.MiddlewareLoggedIn(clilogic.HandlerFollowFeed))
	commands.Register("following", clilogic.MiddlewareLoggedIn(clilogic.HandlerFeedFollowsForUser))
	commands.Register("unfollow", clilogic.MiddlewareLoggedIn(clilogic.HandlerDeleteFeedFollow))
	commands.Register("browse", clilogic.MiddlewareLoggedIn(clilogic.HandlerBrowse))
	// Process the command entered by the user
	runCommandEntered(appState, commands)

	os.Exit(0)
}
