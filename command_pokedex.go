package main

import "fmt"

func commandPokedex(config *Config) error {
	if len(config.pokedex) == 0 {
		fmt.Print("No pokemon currently stored in your Pokedex!")
		return nil
	}
	for name := range config.pokedex {
		fmt.Printf("Your Pokedex:\n")
		fmt.Printf("  - %v\n", name)
	}
	return nil
}
