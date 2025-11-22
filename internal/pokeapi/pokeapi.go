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