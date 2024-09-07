package main

import (
	"github.com/joshparkerj/pokedexcli/internal/pokeapi"
	"strings"
)

func commandExplore(location string) (result string, err error) {
	results, err := pokeapi.Explorer(location)
	if err != nil {
		return
	}

	result = strings.Join(results, "\n")
	return
}
