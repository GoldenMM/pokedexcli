package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/GoldenMM/pokedexcli/internal/pokeapi"
	"github.com/GoldenMM/pokedexcli/internal/pokecache"
)

func main() {
	// Create the scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Print the welcome message
	fmt.Println("Go REPL")
	fmt.Println("Type 'exit' to quit")

	// Create the cache
	cache := pokecache.NewCache(5 * time.Second)

	// Import the commands
	commandMap := getCLICommands()

	// Create the configuration
	config := &Config{
		next:     "https://pokeapi.co/api/v2/location-area/",
		previous: "",
		pokedex:  make(map[string]pokeapi.Pokemon)}

	// Start the REPL and the control loop
	for {
		fmt.Print("pokedex>> ") // line message

		// Check if the scanner has a token
		if !scanner.Scan() {
			break
		}
		// Get the input and check if it had a command or not
		input := strings.Fields(scanner.Text())
		var commandString string
		var p string
		if len(input) == 1 {
			commandString = input[0]
		}
		if len(input) == 2 {
			commandString = input[0]
			p = input[1]
		}
		if len(input) > 2 {
			fmt.Println("Invalid input, too many arguments")
			continue
		}

		// Check if the input is a command
		command, ok := commandMap[commandString]
		if !ok {
			fmt.Println("Command not found")
			continue
		}
		// Execute the command
		err := command.callback(config, p, cache)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
