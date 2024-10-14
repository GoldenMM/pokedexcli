package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCLICommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println("The following commands are available:")
	for _, command := range getCLICommands() {
		fmt.Printf("%s: \t\t %s\n", command.name, command.description)
	}
	return nil
}

func commandExit() error {
	println("Exiting the Pokedex.")
	os.Exit(0)
	return nil
}
