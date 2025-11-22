package pokeapi

import(
	"fmt"
	"net/http"
	"encoding/json"
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

func GetLocationArea(url string, client *http.Client) (LocationArea, error) {

	req, err := http.NewRequest("GET", url, nil)

	var location LocationArea
	if err != nil {
		return location, fmt.Errorf("error creating api request: %s", err)
	}

	res, err := client.Do(req)
	if err !=nil {
		return location, fmt.Errorf("error getting api response: %s", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&location); err != nil {
		return location, fmt.Errorf("error decoding api response: %s", err)
	}

	return location, nil
}