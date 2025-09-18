package main

import (
	"fmt"
)

func commandMap(config *Config) error {
	var url string
	if config.next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"

	} else {
		url = config.next
	}
	data, err := config.pokeapiClient.GetLocationData(url)
	if err != nil {
		return fmt.Errorf("error getting location data: %w", err)
	}
	config.previous = data.Previous
	config.next = data.Next
	locations := data.Results
	for _, location := range locations {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}

func commandMapb(config *Config) error {
	var url string
	if config.next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"

	} else {
		url = config.next
	}
	data, err := config.pokeapiClient.GetLocationData(url)
	if err != nil {
		return fmt.Errorf("error getting location data: %w", err)
	}
	config.previous = data.Previous
	config.next = data.Next
	locations := data.Results
	for _, location := range locations {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}
