package main

import (
	"fmt"
)

func commandExplore(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/" + config.location
	data, err := config.pokeapiClient.GetEncounterablePokemonData(url)
	if err != nil {
		return fmt.Errorf("error getting encounterable pokemon data: %w", err)
	}
	for _, pokemon := range data.PokemonEncounters {
		fmt.Printf("%v\n", pokemon.Pokemon["name"])
	}
	return nil
}
