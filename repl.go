package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/collinmcguire1/pokedex/internal/pokeapi"
)

type config struct {
	commands         map[string]cliCommand
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error checking stdin:", err)
		return
	}
	piped := info.Mode()&os.ModeCharDevice == 0

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type exit to quit!")

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
			return
		}
		input := scanner.Text()
		if piped {
			fmt.Println(input)
		}
		words := cleanInput(input)
		if len(words) == 0 {
			fmt.Println("Oops! An empty string was input. Try typing something before hitting 'enter' to avoid running into a panic crash in a less fool-proof program!")
			continue
		}
		command, exists := cfg.commands[words[0]]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("------------------------")
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a (hopefully) helpful help message (this one!)",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show the name of 20 areas. Each subsequent call of 'map' goes to the next 20.",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "'Map-Back' -- Show the previous 20 areas.",
			callback:    commandMapb,
		},
	}
}
