package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInputSlice := cleanInput(input)
		if len(cleanedInputSlice) == 0 {
			fmt.Println("Oops! An empty string was input. Try typing something before hitting 'enter' to avoid running into a panic crash!")
			continue
		}
		firstWord := cleanedInputSlice[0]
		command, ok := commands[firstWord]
		if ok {
			result := command.callback()
			if result != nil {
				fmt.Println(result)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}

}

var commands = map[string]cliCommand{}

func getCommands() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message (this one!)",
			callback:    commandHelp,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex! Use CTRL+C to quit!\n\n\nUsage:")
	fmt.Println()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	fmt.Println()

	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
