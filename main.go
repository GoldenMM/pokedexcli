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
		command.callback()
	}
}
