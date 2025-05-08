package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (pokemonLocationArea, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		respLocations := pokemonLocationArea{}
		err := json.Unmarshal(val, &respLocations)
		if err != nil {
			return pokemonLocationArea{}, err
		}
		return respLocations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return pokemonLocationArea{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("Error in decoding the response body ", err)
		return pokemonLocationArea{}, err
	}
	locResp := pokemonLocationArea{}
	err = json.Unmarshal(body, &locResp)

	if err != nil {
		fmt.Println(err)
	}
	c.cache.Add(url, body)

	return locResp, nil

}
