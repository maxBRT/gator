package clilogic

// Command represents a CLI command with its name and arguments
// Used to parse and process user input from the command line
type Command struct {
	Name string   // The name of the command (e.g., "login")
	Args []string // Arguments passed to the command
}
