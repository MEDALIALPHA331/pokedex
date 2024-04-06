package main

import (
	"fmt"
	"log"

	pokeapi "github.com/MEDALIALPHA331/pokedex/internal/pokeapi"
)

var (
	config = pokeapi.NewConfig()
)

func commandMap() error {
	locations, err := pokeapi.GetNextPokeLocations(&config)

	if err != nil {
		log.Fatal(err)
	}

	for _, location := range locations.Results {
		fmt.Printf("%+v \n", location.Name)
	}

	return nil
}
