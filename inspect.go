package main

import (
	"fmt"
	"github.com/joshparkerj/pokedexcli/internal/pokedex"
)

func commandInspect(pokemonName string) (result string, err error) {
	pokemon, ok := pokedex.Pokedex.Get(pokemonName)
	if !ok {
		result = "you have not caught that pokemon"
		return
	}

	result = fmt.Sprintf("Name: %v\n", pokemon.Name)
	result += fmt.Sprintf("Height: %v\n", pokemon.Height)
	result += fmt.Sprintf("Weight: %v\n", pokemon.Weight)
	result += "Stats:\n"
	for _, stat := range pokemon.Stats {
		result += fmt.Sprintf(" - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	result += "Types:\n"
	for _, pokemonType := range pokemon.Types {
		result += fmt.Sprintf(" - %v\n", pokemonType.Type.Name)
	}

	return
}
