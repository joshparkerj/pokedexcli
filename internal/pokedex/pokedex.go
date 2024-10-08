package pokedex

import (
	"github.com/joshparkerj/pokedexcli/internal/pokeapi"
	"sync"
)

type pokedex struct {
	mu  *sync.Mutex
	dex map[string]pokeapi.Pokemon
}

func startPokedex() (p pokedex) {
	p = pokedex{
		mu:  &sync.Mutex{},
		dex: make(map[string]pokeapi.Pokemon),
	}

	return
}

func (p *pokedex) Add(pokemon pokeapi.Pokemon) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.dex[pokemon.Name] = pokemon
}

func (p *pokedex) Get(pokemonName string) (pokemon pokeapi.Pokemon, ok bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	pokemon, ok = p.dex[pokemonName]
	return
}

func (p *pokedex) List() (pokemonNames []string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for name, _ := range p.dex {
		pokemonNames = append(pokemonNames, name)
	}

	return
}

var Pokedex = startPokedex()
