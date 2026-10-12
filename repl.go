package main

import (
	//	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	output := strings.Fields(lower)

	return output
}
