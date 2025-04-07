package clilogic

// Command represents a CLI command input.
// It encapsulates the command name and its arguments
// in a structured format for processing.
//
// Example:
//
//	Command{Name: "login", Args: []string{"username"}}
type Command struct {
	Name string   // The name of the command (e.g., "login")
	Args []string // Arguments passed to the command
}
