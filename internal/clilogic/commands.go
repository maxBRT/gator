package clilogic

import "fmt"

// Commands implements the command registry and execution system.
// It provides a flexible way to register and execute CLI commands
// while maintaining clean separation of concerns.
type Commands struct {
	// Map of command names to their handler functions
	// This allows for dynamic dispatch of commands
	cmdList map[string]func(*State, Command) error
}

// Register adds a new command handler to the registry.
// Parameters:
//   - name: The command name to register
//   - f: The handler function for the command
//
// The handler function receives the application state and command arguments.
func (c *Commands) Register(name string, f func(*State, Command) error) {
	// Initialize the map if this is the first registration
	// This ensures we don't need a separate initialization step
	if c.cmdList == nil {
		c.cmdList = make(map[string]func(*State, Command) error)
	}
	c.cmdList[name] = f
}

// Run executes a command if it exists in the registered commands.
// It looks up the command by name and invokes the corresponding handler.
// Returns any error from the command execution.
func (c *Commands) Run(State *State, cmd Command) error {
	if f, ok := c.cmdList[cmd.Name]; ok {
		return f(State, cmd)
	}
	// Provide feedback when a command is not recognized
	fmt.Println("Command not found.")
	return nil
}
