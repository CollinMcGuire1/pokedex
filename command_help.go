package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex! Use CTRL+C to quit!")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range cfg.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	fmt.Println()

	return nil
}
