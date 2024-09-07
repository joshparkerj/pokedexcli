package main

import (
	"fmt"
	"github.com/joshparkerj/pokedexcli/internal/pokeapi"
)

func commandCatch(pokemon string) (result string, err error) {
	success, err := pokeapi.Catch(pokemon)
	if err != nil {
		return
	}

	result = fmt.Sprintf("Throwing a Pokeball at %v...\n", pokemon)
	if success {
		result += fmt.Sprintf("%v was caught!", pokemon)
	} else {
		result += fmt.Sprintf("%v escaped!", pokemon)
	}

	return
}
