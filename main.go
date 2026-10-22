package main

//	TODO:

// 			add more tests for pokecache_test;
// 				- missing key test for Get()
// 				- Add()ing the same key twice - should it overwrite, or give an error, etc.
// 				- Check that old entries are properly being reaped while new entries remain
// 				- concurrent access: Multiple goroutines call Add and Get while reaping runs.
// 					Run go test -race ./... to detect data races.

//			put error checks for invalid locations, pokemon, etc. ex 'explore rock' should 
// 				immediately return with 'location not found'

//			in command_catch.go, add a line below 'fmt.Printf("%s was caught!", pokemon.Name)'
//				which checks if it is in the pokedex already, and print a new line if it is not,
//				which informs the user that this new pokemon is being added.
// 			also modify catch calc to use species CaptureRate rather than pokemon baseXP

import (
	"time"

	"github.com/collinmcguire1/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	startRepl(cfg)
}
