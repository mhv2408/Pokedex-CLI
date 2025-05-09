package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
	"time"
)

func cleanInput(text string) []string {
	//convert the user input into a slice of strings sperated by space...convert every word to lowercase and
	// remove all the leading and trailing spaces
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	config_Ptr := &config{
		pokeapliClient: pokeClient,
		caughtPokemon:  make(map[string]pokeapi.Pokemon),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		input_slice := cleanInput(input)
		if len(input_slice) == 0 {
			continue
		}
		command_name := input_slice[0]
		args := []string{}
		if len(input_slice) > 1 {
			args = input_slice[1:]
		}
		command, exists := getCommands()[command_name]
		if exists {
			err := command.callback(config_Ptr, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown Command")
		}
	}
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
			description: "Displays the names of 20 location areas.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 location areas.",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays list of all Pokemons available in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to Catch the given Pokemon",
			callback:    commandCatch,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapliClient  pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	caughtPokemon   map[string]pokeapi.Pokemon
}
