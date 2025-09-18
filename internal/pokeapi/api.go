package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationsResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type pokemonResponse struct {
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

func (c *Client) GetLocationData(url string) (LocationsResponse, error) {
	//check cache for data
	if entry, ok := c.cache.Get(url); ok {
		data := LocationsResponse{}
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return LocationsResponse{}, fmt.Errorf("error decoding location data from cache: %w", err)
		}
		return data, nil
	}
	//create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsResponse{}, fmt.Errorf("error creating locations request: %w", err)
	}
	//use client to call endpoint
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsResponse{}, fmt.Errorf("error calling location area endpoint: %w", err)
	}
	defer res.Body.Close()
	//read response into []bytes
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsResponse{}, fmt.Errorf("error reading bytes from response: %w", err)
	}
	//umarshal []bytes into response struct
	data := LocationsResponse{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return LocationsResponse{}, fmt.Errorf("error unmarshaling bytes: %w", err)
	}
	c.cache.Add(url, bytes)
	return data, nil
}

func (c *Client) GetPokemonData(url string) (pokemonResponse, error) {
	if entry, ok := c.cache.Get(url); ok {
		data := pokemonResponse{}
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return pokemonResponse{}, fmt.Errorf("error decoding pokemon data from cache: %w", err)
		}
		return data, nil
	}
	//create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemonResponse{}, fmt.Errorf("error creating locations request: %w", err)
	}
	//use client to call endpoint
	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokemonResponse{}, fmt.Errorf("error calling location area endpoint: %w", err)
	}
	defer res.Body.Close()
	//read response into []bytes
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemonResponse{}, fmt.Errorf("error reading bytes from response: %w", err)
	}
	//umarshal []bytes into response struct
	data := pokemonResponse{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return pokemonResponse{}, fmt.Errorf("error unmarshaling bytes: %w", err)
	}
	c.cache.Add(url, bytes)
	return data, nil
}
