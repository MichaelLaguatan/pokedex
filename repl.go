package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MichaelLaguatan/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	pokeapiClient pokeapi.Client
	pokedex       map[string]pokeapi.Pokemon
	location      string
	pokemon       string
	previous      string
	next          string
}

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if command, ok := getCommands()[input[0]]; ok {
			if command.name == "explore" {
				config.location = input[1]
			}
			if command.name == "catch" {
				config.pokemon = input[1]
			}
			if command.name == "pokemon" {
				config.pokemon = input[1]
			}
			err := command.callback(config)
			if err != nil {
				fmt.Printf("Error exiting pokedex: %v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all encounterable pokemon in a location area `explore <location name>`",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon to add it to the user's Pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Lists the information about a pokemon in the user's Pokedex",
			callback:    commandInspect,
		},
	}
}
