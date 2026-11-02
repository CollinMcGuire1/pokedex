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

// 			in command_catch.go modify catch calc to use species CaptureRate rather than pokemon baseXP

//			FROM FINAL LESSON OF UNIT:
//			save pokedex to disk so that a user can keep progress between sessions
//			add party system and exp to level up, evolve
//			improve map/mapb to allow 'left' 'right' etc.
//			different types of catching balls (great ball, ultra ball, etc)

//			package the program so that it can be used outside of a terminal and on different OS's.

import (
	"time"

	"github.com/collinmcguire1/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	startRepl(cfg)
}
