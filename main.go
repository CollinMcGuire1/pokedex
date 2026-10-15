package main

// FINISH REFACTORING SO THAT COMMAND_EXIT, COMMAND_HELP, REPL.GO ARE IN THEIR OWN FILES AND MAIN.GO JUST CONFIGS CFG AND CALLS STARTREPL(CFG)

//import (
//"bufio"
//"fmt"
//"os"
//)

func main() {
	cfg := &config{
		commands: getCommands(),
	}
	startRepl(cfg)
}
