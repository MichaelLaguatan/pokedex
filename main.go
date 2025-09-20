package main

import (
	"time"

	"github.com/MichaelLaguatan/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Second*5)
	config := &Config{pokeapiClient: pokeClient, pokedex: map[string]pokeapi.Pokemon{}}
	startRepl(config)
}
