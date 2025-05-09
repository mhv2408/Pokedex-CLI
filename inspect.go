package main

import "fmt"

func commandInspect(configPtr *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you need to enter a pokemon name")
	}
	pokemonName := args[0]
	if pokemon, ok := configPtr.caughtPokemon[pokemonName]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for i := 0; i < len(pokemon.Stats); i++ {
			fmt.Printf("\t- %s: %d\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
		}
		fmt.Println("Types:")
		for i := 0; i < len(pokemon.Types); i++ {
			fmt.Printf("\t- %s\n", pokemon.Types[i].Type.Name)
		}
		return nil
	}

	return fmt.Errorf("you have not caught that pokemon")
}
