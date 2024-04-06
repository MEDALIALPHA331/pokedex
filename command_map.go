package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	pokeapi "github.com/MEDALIALPHA331/pokedex/internal/pokeapi"
)

const POKE_API_URL = "https://pokeapi.co/api/v2"


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

func commandMap() error {
	_, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	res, err := http.Get(POKE_API_URL + "/location/")

	if err != nil {
		return err
	}

	if res.Status != "200 OK" {
		return errors.New("couldn't Fetch Pokemon Location Data")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var data pokeapi.LocationAreas
	json.Unmarshal(body, &data)

	if err != nil {
		return err
	}

	for _, location := range data.Results {
		fmt.Printf("%+v \n", location.Name)
	}

	return nil
}
