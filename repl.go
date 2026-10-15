package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	commands map[string]cliCommand
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !reader.Scan() {
			if err := reader.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
			return
		}
		input := reader.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			fmt.Println("Oops! An empty string was input. Try typing something before hitting 'enter' to avoid running into a panic crash!")
			continue
		}
		command, exists := cfg.commands[words[0]]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
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
	}
}
