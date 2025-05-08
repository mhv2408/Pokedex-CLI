package main

import (
	"fmt"
)

func commandExplore(configPtr *config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("you must provide a location name")
	}

	areaPokemeonInfo, err := configPtr.pokeapliClient.GetLocation(args[0])

	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", areaPokemeonInfo.Name)
	fmt.Println("Found Pokemon: ")
	for i := 0; i < len(areaPokemeonInfo.PokemonEncounters); i++ {

		fmt.Println(areaPokemeonInfo.PokemonEncounters[i].Pokemon.Name)
	}
	return nil
}
