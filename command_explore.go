package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationResponse struct {
	Id                   int            `json:"id"`
	Name                 string         `json:"name"`
	GameIndex            int            `json:"game_index"`
	EncounterMethodRates []any          `json:"encounter_method_rates"`
	Location             map[string]any `json:"location"`
	Names                []any          `json:"names"`
	PokemonEncounters    []struct {
		Pokemon        map[string]string `json:"pokemon"`
		VersionDetails []any             `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func explore(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/" + config.location
	if entry, ok := config.cache.Get(url); ok {
		data := locationResponse{}
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return fmt.Errorf("error decoding location data from cache: %w", err)
		}
		for _, pokemon := range data.PokemonEncounters {
			fmt.Printf("%v\n", pokemon.Pokemon["name"])
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
	data := locationResponse{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return fmt.Errorf("error unmarshaling bytes: %w", err)
	}
	for _, pokemon := range data.PokemonEncounters {
		fmt.Printf("%v\n", pokemon.Pokemon["name"])
	}
	config.cache.Add(url, bytes)
	return nil
}
