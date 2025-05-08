package main

import "fmt"

func commandMap(configPtr *config, args ...string) error {
	locResp, err := configPtr.pokeapliClient.ListLocations(configPtr.nextLocationURL)
	if err != nil {
		return err
	}
	configPtr.nextLocationURL = locResp.Next
	configPtr.prevLocationURL = locResp.Previous
	for i := 0; i < len(locResp.Results); i++ {
		fmt.Println(locResp.Results[i].Name)
	}

	return nil
}
