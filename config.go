package main

import (
	"errors"

	"github.com/Waterbootdev/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func (c *config) NextListLocationAreas() (*pokeapi.Resource, error) {
	resource, err := c.pokeapiClient.ListLocationAreas(c.nextLocationsURL)
	return &resource, err
}

func (c *config) PreviousListLocationAreas() (*pokeapi.Resource, error) {

	if c.prevLocationsURL == nil {
		return nil, errors.New("you're on the first page")
	}

	resource, err := c.pokeapiClient.ListLocationAreas(c.prevLocationsURL)

	return &resource, err
}
