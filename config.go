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

func (c *config) NextListLocationAreas() (pokeapi.Resource, error) {
	return c.pokeapiClient.ListLocationAreas(c.nextLocationsURL)
}

func (c *config) PreviousListLocationAreas() (pokeapi.Resource, error) {

	if c.prevLocationsURL == nil {
		return pokeapi.Resource{}, errors.New("you're on the first page")
	}

	return c.pokeapiClient.ListLocationAreas(c.prevLocationsURL)
}
