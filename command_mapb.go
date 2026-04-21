package main

import (
	"fmt"

	"github.com/Nico-DM/pokedexcli/internal/pokeapi"
)

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	res, err := api.Get(*cfg.Previous)
	if err != nil {
		return err
	}

	locationAreas, err := api.Unmarshal[api.LocationAreaPage](res)
	if err != nil {
		return err
	}

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	cfg.Next = *locationAreas.Next
	cfg.Previous = locationAreas.Previous

	return nil
}