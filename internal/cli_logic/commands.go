package clilogic

type commands struct {
	// The list of commands
	cmdList map[string]func(*State, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	panic("not implemented")
}
