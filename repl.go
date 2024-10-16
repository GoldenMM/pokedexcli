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
	pokedex  map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string, *pokecache.Cache) error
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
		"explore": {
			name:        "explore",
			description: "Explore the map locations",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display the pokedex",
			callback:    commandPokedex,
		},
	}
}

func commandHelp(config *Config, _ string, _ *pokecache.Cache) error {
	fmt.Println("The following commands are available:")
	fmt.Println("=====================================")
	// TODO: Sort the commands
	for _, command := range getCLICommands() {
		fmt.Printf("%s: \t\t %s\n", command.name, command.description)
	}
	return nil
}

func commandExit(config *Config, _ string, _ *pokecache.Cache) error {
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

func commandCatch(config *Config, p string, cache *pokecache.Cache) error {
	// Check if the input is empty
	if p == "" {
		return fmt.Errorf("no pokemon provided")
	}

	// Check if the pokemon is in the pokedex
	if _, ok := config.pokedex[p]; ok {
		return fmt.Errorf("pokemon already in pokedex")
	}

	var pokemon pokeapi.Pokemon
	// Check if the cache has the data already
	if val, ok := cache.Get(config.next); ok {
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return fmt.Errorf("failed to unmarshal cached data: %v", err)
		}
	} else {
		// Get the pokemon's data from the API wrapper
		var err error
		pokemon, err = pokeapi.GetPokemon(p)
		if err != nil {
			return fmt.Errorf("failed to get map locations: %v", err)
		}

		// Cache the result
		jsonData, err := json.Marshal(pokemon)
		if err != nil {
			return fmt.Errorf("failed to marshal data: %v", err)
		}
		cache.Add(config.next, jsonData)
	}

	// Add the pokemon to the pokedex
	config.pokedex[p] = pokemon
	return nil
}

func commandInspect(config *Config, p string, _ *pokecache.Cache) error {
	// Check if the input is empty
	if p == "" {
		return fmt.Errorf("no pokemon provided")
	}

	// Check if the pokemon is in the pokedex
	pokemon, ok := config.pokedex[p]
	if !ok {
		return fmt.Errorf("pokemon not in pokedex")
	}

	// Print the pokemon's data
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Types: ")
	for _, t := range pokemon.Types {
		fmt.Printf("\t- %s\n", t.Type.Name)
	}
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf("\t- %s\n", ability.Ability.Name)
	}

	return nil
}

func commandPokedex(config *Config, _ string, _ *pokecache.Cache) error {
	// Check if the pokedex is empty
	if len(config.pokedex) == 0 {
		return fmt.Errorf("pokedex is empty")
	}

	// Print the pokedex
	fmt.Println("Your Pokedex:")
	for name := range config.pokedex {
		fmt.Printf("\t- %s\n", name)
	}
	return nil
}
