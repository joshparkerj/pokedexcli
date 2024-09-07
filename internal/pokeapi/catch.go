package pokeapi

import (
	"math/rand"
)

func Catch(pokemonName string) (result bool, pokemon Pokemon, err error) {
	pokemon, err = getPokemon(pokemonName)
	if err != nil {
		return
	}

	baseXp := pokemon.BaseExperience
	captureChance := 0.9*(1.0-(36.0+float64(baseXp))/644.0) + 0.1
	roll := rand.Float64()
	result = roll < captureChance
	return
}
