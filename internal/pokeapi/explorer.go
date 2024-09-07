package pokeapi

func Explorer(locationName string) (result []string, err error) {
	r, err := getLocationPokemon(locationName)
	if err != nil {
		return
	}

	for _, rr := range r.PokemonEncounters {
		result = append(result, rr.Pokemon.Name)
	}

	return
}
