package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const POKE_API_URL = "https://pokeapi.co/api/v2"

func commandMap() error {
	res, err := http.Get(POKE_API_URL + "/location/?limit=20&offset=20")

	if err != nil {
		log.Fatalf("Error getting pokemons from api: %+v", err)
	}

	if res.Status != "200 OK" {
		fmt.Println("Couldn't Fetch Pokemon Location Data")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Couldn't Read body: %+v", err)
	}

	var data LocationResult
	json.Unmarshal(body, &data)

	if err != nil {
		log.Fatalf("%+v", err)
	}

	for _, location := range data.Results {
		fmt.Printf("%+v \n", location.Name)
	}

	return nil
}
