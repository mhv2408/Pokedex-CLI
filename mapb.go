package main

import "fmt"

func commandMapB(configPtr *config, args ...string) error {
	if configPtr.prevLocationURL == nil {
		return fmt.Errorf("you're on the first page")
	}
	respLoc, err := configPtr.pokeapliClient.ListLocations(configPtr.prevLocationURL)
	if err != nil {
		return err
	}

	configPtr.nextLocationURL = respLoc.Next
	configPtr.prevLocationURL = respLoc.Previous

	for i := 0; i < len(respLoc.Results); i++ {
		fmt.Println(respLoc.Results[i].Name)
	}
	return nil
}
