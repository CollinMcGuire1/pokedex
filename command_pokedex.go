package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Here's what's in your Pokedex so far:")
	for _, p := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
