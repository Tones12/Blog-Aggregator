package main

import (
	"github.com/tones12/blog-aggregator/internal/config"
	"fmt"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
    	log.Fatalf("error reading config: %v", err)
	}
	
	newState := state{cfg: &cfg}

	newCommands := commands{commandMap: make(map[string]func(*state, command) error)}

	newCommands.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		fmt.Print("Error, not enough arguments provided\n")
		os.Exit(1)
	}
	commandName := args[1]
	argSlice := args[2:]

	newCommand := command{name: commandName, args: argSlice}

	err = newCommands.run(&newState, newCommand)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}