package main

import (
	"errors"

	"github.com/Waterbootdev/pokedex/internal/pokeapi"
)

func InitConfig(limit int) (config, error) {

	res, err := pokeapi.PokeLocationAreaResource()

	if err != nil {
		return config{}, err
	}

	if limit >= res.Count {
		return config{}, errors.New("limit is to big")
	}

	config := config{
		LocationAreas:        make([]string, res.Count),
		init:                 true,
		limit:                limit,
		length:               res.Count,
		currentLocationAreas: make([]string, limit),
	}

	for i := 0; i < res.Count; i++ {
		name := res.Results[i].Name
		config.LocationAreas[i] = name
	}

	return config, nil
}
