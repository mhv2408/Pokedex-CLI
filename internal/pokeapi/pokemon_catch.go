package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemon_name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon_name

	if val, ok := c.cache.Get(url); ok {
		resPokemon := Pokemon{}
		err := json.Unmarshal(val, &resPokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return resPokemon, nil
	}
	resPokemon := Pokemon{}
	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	val, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return Pokemon{}, err
	}
	err = json.Unmarshal(val, &resPokemon)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, val)
	return resPokemon, nil

}
