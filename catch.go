package main

import (
	"fmt"
	"github.com/joshparkerj/pokedexcli/internal/pokeapi"
	"github.com/joshparkerj/pokedexcli/internal/pokedex"
)

func commandCatch(pokemonName string) (result string, err error) {
	success, pokemon, err := pokeapi.Catch(pokemonName)
	if err != nil {
		return
	}

	result = fmt.Sprintf("Throwing a Pokeball at %v...\n", pokemonName)
	if success {
		result += fmt.Sprintf("%v was caught!", pokemonName)
		pokedex.Pokedex.Add(pokemon)
	} else {
		result += fmt.Sprintf("%v escaped!", pokemonName)
	}

	return
}
