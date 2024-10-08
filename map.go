package main

import (
	"github.com/joshparkerj/pokedexcli/internal/pokeapi"
	"strings"
)

func commandMap(_ string) (result string, err error) {
	results, err := pokeapi.Locations()
	if err != nil {
		return
	}

	result = strings.Join(results, "\n")
	return
}
