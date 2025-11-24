package main

import(
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandExplore(config *config, name string) error {

	var pokemonList pokeapi.PokemonList

	if name == "" {
		return fmt.Errorf("please provide area name")
	}

	fmt.Printf("Exploring %s...\n", name)
	pok, err := pokeapi.GetPokemonList(name, config.client, config.cache)
	if err != nil {
		return fmt.Errorf("error in getting pokemonlist: %s", err)
	}
	pokemonList = pok

	if len(pokemonList.PokemonEncounters) <= 0 {
		return fmt.Errorf("got empty list of pokemon")
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range(pokemonList.PokemonEncounters) {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
