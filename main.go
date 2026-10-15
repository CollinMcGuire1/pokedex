package main

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
