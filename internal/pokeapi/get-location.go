package pokeapi

import (
	"encoding/json"
	"fmt"
)

func getLocation(locationName string) (result Location, err error) {
	path := fmt.Sprintf("location/%v/", locationName)
	body, err := get(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &result)
	return
}
