package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + config.pokemon
	pokemon, err := config.pokeapiClient.GetPokemonData(url)
	if err != nil {
		return fmt.Errorf("error getting pokemon data: %w", err)
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	if catchRate := rand.Float64(); catchRate >= 0.5 {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		if _, ok := config.pokedex[pokemon.Name]; !ok {
			config.pokedex[pokemon.Name] = pokemon.ConvertToPokemon()
		}
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}
	return nil
}
