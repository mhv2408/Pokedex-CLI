package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func explore(configPtr *config, location string) error {

	location_URL := curr_URL + location
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
