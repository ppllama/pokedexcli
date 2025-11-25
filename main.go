package main

import (
	"net/http"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
	"time"
)

type config struct {
	client		*http.Client
	cache		*pokecache.Cache
	pokedex		map[string]pokeapi.Pokemon
	next		*string
	previous	*string
}

func main() {
	conf := config{
		client: pokeapi.NewClient(),
		cache: pokecache.NewCache(time.Minute * 5),
		pokedex: map[string]pokeapi.Pokemon{},
	}

	StartRepl(&conf)
}