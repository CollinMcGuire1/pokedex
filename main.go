package main

// FINISH REFACTORING SO THAT COMMAND_EXIT, COMMAND_HELP, REPL.GO ARE IN THEIR OWN FILES AND MAIN.GO JUST CONFIGS CFG AND CALLS STARTREPL(CFG)

import (
	"bufio"
	"fmt"
	"os"
)

func main(cfg *config) {
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
		command, exists := cfg.commands[firstWord]
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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("\nWelcome to the Pokedex! Use CTRL+C to quit!\n\n\nUsage:")
	fmt.Println()
	for _, command := range config.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	fmt.Println()

	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands map[string]cliCommand
}
