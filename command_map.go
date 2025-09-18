package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type response struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}

type locationArea struct {
	Id                   int            `json:"id"`
	Name                 string         `json:"name"`
	Index                int            `json:"game_index"`
	EncounterMethodRates []any          `json:"encounter_method_rates"`
	Location             map[string]any `json:"location"`
	Names                []any          `json:"names"`
	PokemonEncounters    []any          `json:"pokemon_encounters"`
}

func commandMap(config *Config) error {
	var url string
	if config.next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"

	} else {
		url = config.next
	}
	if entry, ok := config.cache.Get(url); ok {
		data := response{}
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return fmt.Errorf("error decoding location data from cache: %w", err)
		}
		config.previous = data.Previous
		config.next = data.Next
		locations := data.Results
		for _, location := range locations {
			fmt.Printf("%v\n", location.Name)
		}
		return nil
	}
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error calling location area endpoint: %w", err)
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading bytes from response: %w", err)
	}
	data := response{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return fmt.Errorf("error unmarshaling bytes: %w", err)
	}
	config.previous = data.Previous
	config.next = data.Next
	locations := data.Results
	for _, location := range locations {
		fmt.Printf("%v\n", location.Name)
	}
	config.cache.Add(url, bytes)
	return nil
}

func commandMapb(config *Config) error {
	var url string
	if config.previous == "" {
		return fmt.Errorf("error no previous page of location area names")

	} else {
		url = config.previous
	}
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error calling location area endpoint: %w", err)
	}
	defer res.Body.Close()
	data := response{}
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&data); err != nil {
		return fmt.Errorf("error decoding location data: %w", err)
	}
	config.previous = data.Previous
	config.next = data.Next
	locations := data.Results
	for _, location := range locations {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}
