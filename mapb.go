package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapB(configPtr *config, location string) error {

	if url, ok := configPtr.Prev.(string); ok {
		curr_URL = url
	} else {
		fmt.Println("you're on the first page")
		curr_URL = "https://pokeapi.co/api/v2/location-area/"
	}

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
	return nil
}
