package main

import(
	"fmt"
)

func commandInspect(config *config, name string) error {

	if name == "" {
		return fmt.Errorf("please provide pokemon name")
	}

	foundPokemon, ok := config.pokedex[name]
	if !ok {
		fmt.Printf("you have not caught %s\n", name)
		return nil
	}

	fmt.Printf("Name: %s\n", foundPokemon.Name)
	fmt.Printf("Height: %d\n", foundPokemon.Height)
	fmt.Printf("Weight: %d\n", foundPokemon.Weight)

	if len(foundPokemon.Stats) > 0 {
		fmt.Println("Stats:")
		for _, stat := range(foundPokemon.Stats) {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
	}

	if len(foundPokemon.Types) > 0 {
		fmt.Println("Types:")
		for _, tipe := range(foundPokemon.Types) {
			fmt.Printf("  -%s\n", tipe.Type.Name)
		}
	}

	return nil
}