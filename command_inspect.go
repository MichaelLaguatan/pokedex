package main

import "fmt"

func commandInspect(config *Config) error {
	if entry, ok := config.pokedex[config.pokemon]; ok {
		fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\nStats:\n", entry.Name, entry.Height, entry.Weight)
		for name, value := range entry.Stats {
			fmt.Printf("  -%v: %v\n", name, value)
		}
		fmt.Printf("Types:\n")
		for _, name := range entry.Types {
			fmt.Printf("  - %v\n", name)
		}
	} else {
		fmt.Printf("%v is not in your Pokedex", config.pokemon)
		return fmt.Errorf("%v is not in the Pokedex", config.pokemon)
	}
	return nil
}
