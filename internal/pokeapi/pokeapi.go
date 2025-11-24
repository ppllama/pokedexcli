package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
)


type LocationArea struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func NewClient() *http.Client {
	return &http.Client{}
}

func GetLocationArea(url string, client *http.Client, cache *pokecache.Cache) (LocationArea, error) {

	var location LocationArea

	data, ok := cache.Get(url)
	if ok {
		if err := json.Unmarshal(data, &location); err != nil {
			return location, fmt.Errorf("error getting location cache")
		}
		fmt.Println("locations retrieved from cache")
		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return location, fmt.Errorf("error creating api request: %s", err)
	}

	res, err := client.Do(req)
	if err !=nil {
		return location, fmt.Errorf("error getting api response: %s", err)
	}

	defer res.Body.Close()

	newData, err := io.ReadAll(res.Body)
	if err != nil {
		return location, fmt.Errorf("error reading response: %s", err)
	}

	if err := json.Unmarshal(newData, &location); err != nil {
		return location, fmt.Errorf("error decoding api response %s", err)
	}

	cache.Add(url, newData)
	
	return location, nil
}

type PokemonList struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetPokemonList(area string, client *http.Client, cache *pokecache.Cache) (PokemonList, error) {
	var pokemonList PokemonList

	url := "https://pokeapi.co/api/v2/location-area/" + area
	fmt.Println(url)

	data, ok := cache.Get(url)
	if ok {
		if err := json.Unmarshal(data, &pokemonList); err != nil {
			return pokemonList, fmt.Errorf("error getting pokemonlist cache")
		}
		fmt.Println("pokemonlist retrieved from cache")
		return pokemonList, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return pokemonList, fmt.Errorf("error creating api request: %s", err)
	}

	res, err := client.Do(req)
	if err !=nil {
		return pokemonList, fmt.Errorf("error getting api response: %s", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return pokemonList, fmt.Errorf("not found")
	}

	newData, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemonList, fmt.Errorf("error reading response: %s", err)
	}

	if err := json.Unmarshal(newData, &pokemonList); err != nil {
		return pokemonList, fmt.Errorf("error decoding api response %s", err)
	}

	cache.Add(url, newData)
	
	return pokemonList, nil
}