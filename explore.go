package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func explore(configPtr *config, location string) error {

	location_URL := curr_URL + location
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
	loc := pokemonLocation{}
	err = json.Unmarshal(body, &loc)

	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(loc.Results); i++ {
		fmt.Println(loc.Results[i].Name)
	}
	return nil
}
