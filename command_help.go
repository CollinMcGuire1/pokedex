package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	// Check repl.go for an accurate list of available commands, and update this list here:

	commands := getCommands()
	orderedCommands := []string{
		"help",
		"exit",
		"map",
		"mapb",
		"explore",
		"catch",
		"inspect",
		"pokedex",
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("Welcome to the Pokedex! Type 'exit' to quit!")
	fmt.Println("Commands are case insensitive.")
	fmt.Println()
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range orderedCommands {
		cmd, ok := commands[command]
		if ok {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}
	}
	fmt.Println()
	fmt.Println()

	return nil
}
