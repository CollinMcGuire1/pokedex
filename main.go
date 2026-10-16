package main

//	TODO:
//			add aliases, like 'mapBack', etc
//			add case-insensitivity for commands

import (
	"time"

	"github.com/collinmcguire1/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
