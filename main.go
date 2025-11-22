package main

import (
	"net/http"
	"pokedexcli/internal/pokeapi"
)

type config struct {
	client		*http.Client
	next		*string
	previous	*string
}

func main() {
	conf := config{
		client: pokeapi.NewClient(),
	}

	StartRepl(&conf)
}