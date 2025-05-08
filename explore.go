package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(configPtr *config, location ...string) error {

	if len(location) != 1 {
		return fmt.Errorf("There cannot be more than 1 location...")
	}

	location_URL := curr_URL + location[0]
	fmt.Println(location_URL)
	res, err := http.Get(location_URL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("Error in decoding the response body ", err)
		return err
	}
	location_area := pokemonLocationArea{}
	err = json.Unmarshal(body, &location_area)

	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(location_area.PokemonEncounters); i++ {

		fmt.Println(location_area.PokemonEncounters[i].Pokemon.Name)
	}
	return nil
}
