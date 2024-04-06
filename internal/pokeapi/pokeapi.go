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

const (
	POKE_API_URL = "https://pokeapi.co/api/v2/location"
)

type Config struct {
	Next     string
	Previous string
}

func NewConfig() Config {
	return Config{
		Next:     "https://pokeapi.co/api/v2/location",
		Previous: "",
	}
}

type httpClient struct {
	client *http.Client
}

func NewHttpClient() *httpClient {
	return &httpClient{
		client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func GetNextPokeLocations(config *Config) (*LocationAreas, error) {
	result, err := http.Get(config.Next)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)

	if err != nil {
		return nil, err
	}
	var data LocationAreas
	json.Unmarshal(body, &data)

	config.Next = data.Next
	config.Previous = data.Previous

	return &data, nil
}
