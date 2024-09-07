package main

import (
	"github.com/joshparkerj/pokedexcli/internal/pokeapi"
	"strings"
)

func commandMapb(_ string) (result string, err error) {
	results, err := pokeapi.LocationsBack()
	if err != nil {
		return
	}

	result = strings.Join(results, "\n")
	return
}
