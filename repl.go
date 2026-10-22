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
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error checking stdin:", err)
		return
	}
	piped := info.Mode()&os.ModeCharDevice == 0

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Type 'help' for.... well...")
	fmt.Println("Type 'exit' to quit!")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - -")

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
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		command, exists := cfg.commands[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - -")
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

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	mapCommandf := cliCommand{
		name:        "map, mapf, mapforward",
		description: "Show the name of 20 areas. Each subsequent call of 'map' goes to the next 20.",
		callback:    commandMapf,
	}

	mapCommandb := cliCommand{
		name:        "mapb, mapback",
		description: "'Map Back'. Show the previous 20 areas.",
		callback:    commandMapb,
	}
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

		"map":        mapCommandf,
		"mapf":       mapCommandf,
		"mapforward": mapCommandf,
		"mapb":       mapCommandb,
		"mapback":    mapCommandb,

		"explore": {
			name:        "explore <location_name>",
			description: "Show what pokemon are found in an input location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Show stats about a pokemon in your pokedex",
			callback:    commandInspect,
		},
	}
}
