package main

import (
	"fmt"
)

func explore(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/" + config.location
	data, err := config.pokeapiClient.GetPokemonData(url)
	if err != nil {
		return fmt.Errorf("error unmarshaling bytes: %w", err)
	}
	for _, pokemon := range data.PokemonEncounters {
		fmt.Printf("%v\n", pokemon.Pokemon["name"])
	}
	return nil
}
