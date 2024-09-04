package pokeapi

import (
	"encoding/json"
	"fmt"
)

func getLocationArea(pageNumber int) (result locationArea, err error) {
	path := fmt.Sprintf("location-area?offset=%v&limit=20", (pageNumber-1)*20)
	body, err := get(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &result)
	return
}
