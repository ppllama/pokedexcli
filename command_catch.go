package main

import(
	"fmt"
	"pokedexcli/internal/pokeapi"
	"math/rand"
)

func commandCatch(config *config, name string) error {

	var pokemonData pokeapi.Pokemon

	if name == "" {
		return fmt.Errorf("please provide pokemon name")
	}

	pok, err := pokeapi.GetPokemon(name, config.client, config.cache)
	if err != nil {
		return fmt.Errorf("error in getting pokemonData: %s", err)
	}
	
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemonData = pok

	luck := rand.Intn(pokemonData.BaseExperience / 20)

	if  luck <= 3 {
		fmt.Printf("%s was caught with luck %d!\n", name, luck)
		if _, ok := config.pokedex[name]; ok {
			fmt.Printf("you already have %s in your pokedex. caught %s was released.\n", name, name)
			return nil
		}
		config.pokedex[name] = pokemonData
	} else {
		fmt.Printf("%s escaped with luck %d!\n", name, luck)
	}

	return nil
}
