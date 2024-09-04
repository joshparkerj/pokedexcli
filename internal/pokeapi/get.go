package pokeapi

import (
	"io"
	"net/http"
)

const base = "https://pokeapi.co/api/v2/"

func get(path string) (body []byte, err error) {
	url := base + path
	res, err := http.Get(url)
	if err != nil {
		return
	}

	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	return
}
