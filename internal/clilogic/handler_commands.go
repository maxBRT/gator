package clilogic

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/maxBRT/gator/internal/database"
	"github.com/maxBRT/gator/internal/rss"
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

func HandlerAggregate(state *State, cmd Command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

// HandlerAddFeed creates a new feed entry and automatically follows it
// for the current user. It requires a feed name and URL as arguments.
func HandlerAddFeed(state *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		fmt.Println("Missing arguments")
		os.Exit(1)
	}

	feed, err := state.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	_, err = state.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	return nil
}

// HandlerFollowFeed allows a user to follow an existing feed by its URL.
// It creates a feed follow relationship between the user and the feed.
func HandlerFollowFeed(state *State, cmd Command, user database.User) error {

	if len(cmd.Args) < 1 {
		fmt.Println("Missing arguments")
		os.Exit(1)
	}

	feed, err := state.DB.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Println("That feed does not exist.")
		os.Exit(1)
	}
	_, err = state.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("User %v now follows the %v feed", user.Name, feed.Name)
	return nil
}

// HandlerFeedFollowsForUser retrieves and displays all feeds that
// the current user is following.
func HandlerFeedFollowsForUser(state *State, cmd Command, user database.User) error {
	feeds, err := state.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	if len(feeds) < 1 {
		fmt.Println("User is not following any feeds.")
		os.Exit(0)
	}

	fmt.Printf("User %v is following :\n", user.Name)
	for _, feed := range feeds {
		currentFeed, err := state.DB.GetFeedById(context.Background(), feed.FeedID)
		if err != nil {
			return err
		}
		fmt.Printf("* %v \n", currentFeed.Name)
	}
	return nil
}

// HandlerGetFeeds lists all available feeds in the system
// along with their URLs and creator information.
func HandlerGetFeeds(state *State, cmd Command) error {
	feeds, err := state.DB.GetFeed(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		usr, err := state.DB.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(usr.Name)
	}
	return nil
}

// HandlerDeleteFeedFollow removes a feed follow relationship
// between the current user and a specified feed URL.
func HandlerDeleteFeedFollow(state *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		fmt.Println("Please provide a URL.")
		os.Exit(1)
	}
	feed, err := state.DB.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	err = state.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Feed %v succesfully deleted for user %v", feed.Name, user.Name)
	return nil
}
