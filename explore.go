package main

import (
	"fmt"
	"github.com/joshparkerj/pokedexcli/internal/pokeapi"
)

func commandExplore(location string) (result string, err error) {
	results, err := pokeapi.Explorer(location)
	if err != nil {
		return
	}

	result = fmt.Sprintf("Exploring %v...\nFound Pokemon:", location)

	for _, r := range results {
		result += fmt.Sprintf("\n - %v", r)
	}

	return
}
