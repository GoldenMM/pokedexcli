package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create the scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Print the welcome message
	fmt.Println("Go REPL")
	fmt.Println("Type 'exit' to quit")

	// Import the commands
	commandMap := getCLICommands()

	// Create the configuration
	config := &Config{next: "https://pokeapi.co/api/v2/location-area/", previous: ""}

	// Start the REPL and the control loop
	for {
		fmt.Print("pokedex>> ") // line message

		// Check if the scanner has a token
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		// Check if the input is a command
		command, ok := commandMap[input]
		if !ok {
			fmt.Println("Command not found")
			continue
		}
		// Execute the command
		err := command.callback(config)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
