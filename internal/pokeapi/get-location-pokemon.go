package pokeapi

import (
	"encoding/json"
	"fmt"
)

func getLocationPokemon(locationName string) (result LocationPokemon, err error) {
	path := fmt.Sprintf("location-area/%v/", locationName)
	body, err := get(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &result)
	return
}
