package main

import (
	"bufio"
	"fmt"
	"internal/pokecache"
	"os"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	cache    pokecache.Cache
	previous string
	next     string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &Config{cache: pokecache.NewCache(time.Second * 5)}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if command, ok := getCommands()[input[0]]; ok {
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
	}
}
