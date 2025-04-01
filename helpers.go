package main

import (
	"fmt"
	"os"

	"github.com/maxBRT/gator/internal/clilogic"
	"github.com/maxBRT/gator/internal/config"
)

func checkArgs() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
}

func initState() *clilogic.State {
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	state := &clilogic.State{
		Config: &cfg,
	}
	return state
}

func runCommandEntered(appState *clilogic.State, commands *clilogic.Commands) {
	cmdEntered := clilogic.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err := commands.Run(appState, cmdEntered)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
