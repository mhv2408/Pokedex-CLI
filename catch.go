package main

import (
	"fmt"
	"math/rand"
)

func can_catch(n int) bool {
	p := 1.0 / float64(n)
	return rand.Float64() < p
}

func commandCatch(configPtr *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must have a pokemon to catch")
	}
	pokemonInfo, err := configPtr.pokeapliClient.CatchPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonInfo.Name)
	if can_catch(pokemonInfo.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemonInfo.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonInfo.Name)
	}
	configPtr.caughtPokemon[pokemonInfo.Name] = pokemonInfo
	return nil
}
