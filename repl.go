package main

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"
	"github.com/collinmcguire1/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	rl, err := readline.New("Pokedex > ")
	if err != nil {
		fmt.Println("Error initializing readline:", err)
		return
	}
	defer rl.Close()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Type 'help' for a list of (case insensitive) commands.")
	fmt.Println("Type 'exit' to quit!\n")

	for {
		input, err := rl.Readline()
		if err != nil {
			return
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
		command, exists := getCommands()[commandName]
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
	//in command_help.go, you can add more functions to this list to get an updated print of available commands
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
			description: "Show what pokemon are found in an input <location_written_in_snake_case>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Try to catch a pokemon named <pokemon>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Show stats about a <pokemon> in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See all of the pokemon that you've caught",
			callback:    commandPokedex,
		},
	}
}
