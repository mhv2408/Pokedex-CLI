package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var curr_URL = "https://pokeapi.co/api/v2/location-area/"

func commandMap(configPtr *config) error {

	res, err := http.Get(curr_URL)
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
	loc := pokemonLocation{}
	err = json.Unmarshal(body, &loc)

	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(loc.Results); i++ {
		fmt.Println(loc.Results[i].Name)
	}
	configPtr.Next = loc.Next
	configPtr.Prev = loc.Previous
	curr_URL = loc.Next
	return nil
}
