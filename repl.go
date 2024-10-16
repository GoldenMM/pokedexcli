package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/GoldenMM/pokedexcli/internal/pokeapi"
	"github.com/GoldenMM/pokedexcli/internal/pokecache"
)

type Config struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
}

func getCLICommands(cache *pokecache.Cache) map[string]cliCommand {
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
			callback: func(config *Config, p string) error {
				return commandMap(config, p, cache)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 map locations",
			callback: func(config *Config, p string) error {
				return commandMapb(config, p, cache)
			},
		},
		"explore": {
			name:        "explore",
			description: "Explore the map locations",
			callback: func(config *Config, p string) error {
				return commandExplore(config, p, cache)
			},
		},
	}
}

func commandHelp(config *Config, _ string) error {
	fmt.Println("The following commands are available:")
	fmt.Println("=====================================")
	// TODO: Sort the commands
	for _, command := range getCLICommands(&pokecache.Cache{}) {
		fmt.Printf("%s: \t\t %s\n", command.name, command.description)
	}
	return nil
}

func commandExit(config *Config, _ string) error {
	println("Exiting the Pokedex.")
	os.Exit(0)
	return nil
}

func commandMap(config *Config, _ string, cache *pokecache.Cache) error {
	// Check if the next location-areas exist
	if config.next == "" {
		return fmt.Errorf("no next location-areas")
	}

	var mapLocationRes pokeapi.LocationAreaResp

	// Check if the cache has the next location-areas
	if val, ok := cache.Get(config.next); ok {
		err := json.Unmarshal(val, &mapLocationRes)
		if err != nil {
			return fmt.Errorf("failed to unmarshal cached data: %v", err)
		}
	} else {
		// Get the map locations from the API wrapper
		var err error
		mapLocationRes, err = pokeapi.GetMapLocations(config.next)
		if err != nil {
			return fmt.Errorf("failed to get map locations: %v", err)
		}

		// Cache the result
		jsonData, err := json.Marshal(mapLocationRes)
		if err != nil {
			return fmt.Errorf("failed to marshal data: %v", err)
		}
		cache.Add(config.next, jsonData)
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

func commandMapb(config *Config, _ string, cache *pokecache.Cache) error {
	// Check if the previous location-areas exist
	if config.previous == "" {
		return fmt.Errorf("no previous location-areas")
	}

	var mapLocationRes pokeapi.LocationAreaResp

	// Check if the cache has the next location-areas
	if val, ok := cache.Get(config.next); ok {
		err := json.Unmarshal(val, &mapLocationRes)
		if err != nil {
			return fmt.Errorf("failed to unmarshal cached data: %v", err)
		}
	} else {
		// Get the map locations from the API wrapper
		var err error
		mapLocationRes, err = pokeapi.GetMapLocations(config.next)
		if err != nil {
			return fmt.Errorf("failed to get map locations: %v", err)
		}

		// Cache the result
		jsonData, err := json.Marshal(mapLocationRes)
		if err != nil {
			return fmt.Errorf("failed to marshal data: %v", err)
		}
		cache.Add(config.next, jsonData)
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

func commandExplore(config *Config, p string, cache *pokecache.Cache) error {
	// Check if the input is empty
	if p == "" {
		return fmt.Errorf("no location provided")
	}

	var pokemonInAreaResp pokeapi.PokemonInAreaResp

	// Check if the cache has the data already
	if val, ok := cache.Get(config.next); ok {
		err := json.Unmarshal(val, &pokemonInAreaResp)
		if err != nil {
			return fmt.Errorf("failed to unmarshal cached data: %v", err)
		}
	} else {
		// Get the map locations from the API wrapper
		var err error
		pokemonInAreaResp, err = pokeapi.GetPokemonInArea(p)
		if err != nil {
			return fmt.Errorf("failed to get map locations: %v", err)
		}

		// Cache the result
		jsonData, err := json.Marshal(pokemonInAreaResp)
		if err != nil {
			return fmt.Errorf("failed to marshal data: %v", err)
		}
		cache.Add(config.next, jsonData)
	}

	// Print the pokemon in the area
	for _, pokemon := range pokemonInAreaResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
