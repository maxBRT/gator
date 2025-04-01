package clilogic

import "fmt"

// Commands manages the set of available CLI commands
// It provides registration and execution functionality
type Commands struct {
	// Map of command names to their handler functions
	cmdList map[string]func(*State, Command) error
}

// Register adds a new command to the command list
// Takes a command name and its corresponding handler function
func (c *Commands) Register(name string, f func(*State, Command) error) {
	// Initialize the map if this is the first registration
	if c.cmdList == nil {
		c.cmdList = make(map[string]func(*State, Command) error)
	}
	c.cmdList[name] = f
}

// Run executes a command if it exists in the registered commands
// Returns any error from the command execution
func (c *Commands) Run(State *State, cmd Command) error {
	if f, ok := c.cmdList[cmd.Name]; ok {
		return f(State, cmd)
	}
	fmt.Println("Command not found.")
	return nil
}
