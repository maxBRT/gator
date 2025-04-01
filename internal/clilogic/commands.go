package clilogic

import "fmt"

type Commands struct {
	// The list of commands
	cmdList map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if c.cmdList == nil {
		c.cmdList = make(map[string]func(*State, Command) error)
	}
	c.cmdList[name] = f
}

func (c *Commands) Run(State *State, cmd Command) error {
	if f, ok := c.cmdList[cmd.Name]; ok {
		return f(State, cmd)
	}
	fmt.Println("Command not found.")
	return nil
}
