package main

import (
	"errors"
	"strings"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a single Pokemon name")
	}

	name := strings.ToLower(args[0])
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("You don't have this pokemon in your pokedex yet!")
	}
	height := float64(pokemon.Height) / 10
	weight := float64(pokemon.Weight) / 10
	
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %.2f m\n", height)
	fmt.Printf("Weight: %.1f kg\n", weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		statName := capitalize(stat.Stat.Name)
		fmt.Printf("  -%s: %v\n", statName, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		typeName := capitalize(typeInfo.Type.Name)
		fmt.Println("  -", typeName)
	}

	return nil

}