package main

//	TODO:
//			add aliases, like 'mapBack', etc
//			add case-insensitivity for commands

// 			add more tests for pokecache_test;
// 				- missing key test for Get()
// 				- Add()ing the same key twice - should it overwrite, or give an error, etc.
// 				- Check that old entries are properly being reaped while new entries remain
// 				- concurrent access: Multiple goroutines call Add and Get while reaping runs.
// 					Run go test -race ./... to detect data races.

import (
	"time"

	"github.com/collinmcguire1/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
