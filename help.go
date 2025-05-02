package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	for _, val := range getCommands() {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	return nil
}
