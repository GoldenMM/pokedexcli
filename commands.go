package main

import (
	"fmt"
	"os"

	"github.com/GoldenMM/pokedexcli/internal/pokeapi"
)

type Config struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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
		"map": {
			name:        "map",
			description: "Display next 20 map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 map locations",
			callback:    commandMapb,
		},
	}
}

func commandHelp(config *Config) error {
	fmt.Println("The following commands are available:")
	// TODO: Sort the commands
	for _, command := range getCLICommands() {
		fmt.Printf("%s: \t\t %s\n", command.name, command.description)
	}
	return nil
}

func commandExit(config *Config) error {
	println("Exiting the Pokedex.")
	os.Exit(0)
	return nil
}

func commandMap(config *Config) error {
	// Check if the next location-areas exist
	if config.next == "" {
		return fmt.Errorf("no next location-areas")
	}

	// Get the map locations from the api wrapper
	mapLocationRes, err := pokeapi.GetMapLocations(config.next)
	if err != nil {
		return fmt.Errorf("failed to get map locations: %v", err)
	}

	// Print the map locations
	for _, loc := range mapLocationRes.Results {
		fmt.Println(loc.Name)
	}

	// Update the config values
	config.next = mapLocationRes.Next
	config.previous = mapLocationRes.Previous

	return nil
}

func commandMapb(config *Config) error {
	// Check if the previous location-areas exist
	if config.previous == "" {
		return fmt.Errorf("no previous location-areas")
	}
	// Get the map locations from the api wrapper
	mapLocationRes, err := pokeapi.GetMapLocations(config.previous)
	if err != nil {
		return fmt.Errorf("failed to get map locations: %v", err)
	}

	// Print the map locations
	for _, loc := range mapLocationRes.Results {
		fmt.Println(loc.Name)
	}

	// Update the config values
	config.next = mapLocationRes.Next
	config.previous = mapLocationRes.Previous

	return nil
}
