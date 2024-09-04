package main

import (
	"fmt"
)

func commandHelp() (result string, err error) {
	result = "\nWelcome to the Pokedex!\nUsage:\n\n"
	for _, command := range commands() {
		result += fmt.Sprintf("%v: %v\n", command.name, command.description)
	}

	return
}
