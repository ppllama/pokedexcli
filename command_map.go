package main

import(
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandMap(config *config) error {

	var location pokeapi.LocationArea
	if config.next == nil {
		loc, err := pokeapi.GetLocationArea("https://pokeapi.co/api/v2/location-area?offset=0&limit=20", config.client, config.cache)
		if err != nil {
			return fmt.Errorf("error in getting location: %s", err)
		}
		location = loc
	} else {
		loc, err := pokeapi.GetLocationArea(*config.next, config.client, config.cache)
		if err != nil {
			return fmt.Errorf("error in getting location %s", err)
		}
		location = loc
	}
	
	config.next = location.Next
	config.previous = location.Previous

	if len(location.Results) <= 0 {
		return fmt.Errorf("got empty list of areas")
	}

	for _, area := range(location.Results) {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(config *config) error {

	var location pokeapi.LocationArea
	if config.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		loc, err := pokeapi.GetLocationArea(*config.previous, config.client, config.cache)
		if err != nil {
			return fmt.Errorf("error in getting location %s", err)
		}
		location = loc
	}
	
	config.next = location.Next
	config.previous = location.Previous

	if len(location.Results) <= 0 {
		return fmt.Errorf("got empty list of areas")
	}

	for _, area := range(location.Results) {
		fmt.Println(area.Name)
	}

	return nil
}