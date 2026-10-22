package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a single Pokemon name")
	}

	name := strings.ToLower(args[0])
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	pokeName := capitalize(pokemon.Name)
	//DEBUG LINE	fmt.Printf("%s", pokemon.BaseExperience)
	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)
	if res > 40 {
		fmt.Printf("%s escaped the Pokeball!\n", pokeName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokeName)
	// add a check here to print 'adding <p> to pokedex'

	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}

func capitalize(pokemonName string) string {
	if pokemonName == "" {
		return pokemonName
	}
	return strings.ToUpper(pokemonName[:1]) + pokemonName[1:]
}
