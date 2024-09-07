package main

import (
	"fmt"
	"github.com/joshparkerj/pokedexcli/internal/pokedex"
)

func commandPokedex(_ string) (result string, err error) {
	pokemonNames := pokedex.Pokedex.List()
	result = "Your Pokedex:"
	for _, name := range pokemonNames {
		result += fmt.Sprintf("\n - %v", name)
	}

	return
}
