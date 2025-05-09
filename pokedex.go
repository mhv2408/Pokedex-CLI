package main

import "fmt"

func commandPokedex(configPtr *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range configPtr.caughtPokemon {
		fmt.Printf("\t- %s\n", key)
	}
	return nil
}
