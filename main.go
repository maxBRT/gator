package main

import (
	"fmt"

	"github.com/maxBRT/gator/internal/clilogic"
)

func main() {
	checkArgs()
	appState := initState()

	commands := &clilogic.Commands{}
	commands.Register("login", clilogic.HandlerLogin)

	runCommandEntered(appState, commands)

	// Print the config
	fmt.Println("DB URL:", appState.Config.DBURL)
	fmt.Println("Username:", appState.Config.USERNAME)
}
