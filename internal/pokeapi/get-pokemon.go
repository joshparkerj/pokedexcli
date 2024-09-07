package pokeapi

import (
	"encoding/json"
	"fmt"
)

func getPokemon(pokemon string) (result Pokemon, err error) {
	path := fmt.Sprintf("pokemon/%v/", pokemon)
	body, err := get(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &result)
	return
}
