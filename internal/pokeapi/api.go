package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BatchLocationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationResponse struct {
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

type pokemonResponse struct {
	Id                     int            `json:"id"`
	Name                   string         `json:"name"`
	BaseExperience         int            `json:"base_experience"`
	Height                 int            `json:"height"`
	IsDefault              bool           `json:"is_default"`
	Order                  int            `json:"order"`
	Weight                 int            `json:"weight"`
	Abilities              []any          `json:"abilities"`
	Forms                  []any          `json:"forms"`
	GameIndices            []any          `json:"game_indicies"`
	HeldItems              []any          `json:"held_items"`
	LocationAreaEncounters string         `json:"location_area_encounters"`
	Moves                  []any          `json:"moves"`
	Species                map[string]any `json:"species"`
	Sprites                map[string]any `json:"sprites"`
	Cries                  map[string]any `json:"cries"`
	Stats                  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes     []any `json:"past_types"`
	PastAbilities []any `json:"past_abilities"`
}

type Pokemon struct {
	Name   string
	Weight int
	Height int
	Stats  map[string]int
	Types  []string
}

func (p pokemonResponse) ConvertToPokemon() Pokemon {
	stats := map[string]int{}
	types := []string{}
	for _, stat := range p.Stats {
		stats[stat.Stat.Name] = stat.BaseStat
	}
	for _, entry := range p.Types {
		types = append(types, entry.Type.Name)
	}
	return Pokemon{
		Name:   p.Name,
		Weight: p.Weight,
		Height: p.Height,
		Stats:  stats,
		Types:  types,
	}
}

func (c *Client) GetLocationData(url string) (BatchLocationResponse, error) {
	//check cache for data
	if entry, ok := c.cache.Get(url); ok {
		data := BatchLocationResponse{}
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return BatchLocationResponse{}, fmt.Errorf("error decoding location data from cache: %w", err)
		}
		return data, nil
	}
	//create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return BatchLocationResponse{}, fmt.Errorf("error creating locations request: %w", err)
	}
	//use client to call endpoint
	res, err := c.httpClient.Do(req)
	if err != nil {
		return BatchLocationResponse{}, fmt.Errorf("error calling location area endpoint: %w", err)
	}
	defer res.Body.Close()
	//read response into []bytes
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return BatchLocationResponse{}, fmt.Errorf("error reading bytes from response: %w", err)
	}
	//umarshal []bytes into response struct
	data := BatchLocationResponse{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return BatchLocationResponse{}, fmt.Errorf("error unmarshaling bytes: %w", err)
	}
	c.cache.Add(url, bytes)
	return data, nil
}

func (c *Client) GetEncounterablePokemonData(url string) (LocationResponse, error) {
	if entry, ok := c.cache.Get(url); ok {
		data := LocationResponse{}
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return LocationResponse{}, fmt.Errorf("error decoding pokemon data from cache: %w", err)
		}
		return data, nil
	}
	//create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error creating locations request: %w", err)
	}
	//use client to call endpoint
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error calling location area endpoint: %w", err)
	}
	defer res.Body.Close()
	//read response into []bytes
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error reading bytes from response: %w", err)
	}
	//umarshal []bytes into response struct
	data := LocationResponse{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error unmarshaling bytes: %w", err)
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
