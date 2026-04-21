package main

import (
	"fmt"

	"github.com/Nico-DM/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	res, err := api.Get(cfg.Next)
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