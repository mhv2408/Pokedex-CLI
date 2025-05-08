package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (pokemonLocation, error) {
	url := baseURL + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}
	if val, ok := c.cache.Get(url); ok {
		respLocations := pokemonLocation{}
		err := json.Unmarshal(val, &respLocations)
		if err != nil {
			return pokemonLocation{}, err
		}
		return respLocations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return pokemonLocation{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("Error in decoding the response body ", err)
		return pokemonLocation{}, err
	}
	locResp := pokemonLocation{}
	err = json.Unmarshal(body, &locResp)

	if err != nil {
		fmt.Println(err)
	}
	c.cache.Add(url, body)

	return locResp, nil

}
