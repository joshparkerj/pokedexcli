package pokeapi

import (
	"github.com/joshparkerj/pokedexcli/internal/pokecache"
	"io"
	"net/http"
	"time"
)

const base = "https://pokeapi.co/api/v2/"

var cache = pokecache.NewCache(time.Minute * 5)

func get(path string) (body []byte, err error) {
	body, ok := cache.Get(path)
	if ok {
		return
	}

	url := base + path
	res, err := http.Get(url)
	if err != nil {
		return
	}

	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	cache.Add(path, body)
	return
}
