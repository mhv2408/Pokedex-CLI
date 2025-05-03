package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	//convert the user input into a slice of strings sperated by space...convert every word to lowercase and
	// remove all the leading and trailing spaces
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	config_Ptr := config{}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		input_slice := cleanInput(input)
		first_word := input_slice[0]
		command, exists := getCommands()[first_word]
		if exists {
			err := command.callback(&config_Ptr)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Invalid Command")
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
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type pokemonLocation struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
type config struct {
	Next string
	Prev any
}
