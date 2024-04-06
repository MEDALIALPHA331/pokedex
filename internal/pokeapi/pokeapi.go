package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var (
	POKE_API_URL = "https://pokeapi.co/api/v2/location"
	Client       = NewHttpClient()
)

type Config struct {
	Next     string
	Previous string
}

func NewConfig() Config {
	return Config{
		Next:     POKE_API_URL,
		Previous: "",
	}
}

func NewHttpClient() *http.Client {
	var Client = &http.Client{
		Timeout: time.Minute,
	}
	return Client
}

func GetPokeLocations(config *Config, next bool) (*LocationAreas, error) {

	URL := config.Next
	//! this is not intended behavior
	if !next && len(config.Previous) > 0 {
		//TODO: add isValid function for these urls
		URL = config.Previous
	}

	result, err := Client.Get(URL)

	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)

	if err != nil {
		return nil, err
	}
	var data LocationAreas
	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	config.Next = data.Next
	config.Previous = data.Previous

	return &data, nil
}
