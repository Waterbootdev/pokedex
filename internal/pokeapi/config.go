package pokeapi

import (
	"errors"
	"time"
)

type Config struct {
	pokeapiClient    Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func InitConfig() Config {
	pokeClient := NewClient(5 * time.Second)
	return Config{pokeapiClient: pokeClient}
}

func (c *Config) printNames(locationURL *string) error {

	resource, err := c.pokeapiClient.ListLocationAreas(locationURL)

	if err == nil {
		c.nextLocationsURL = resource.Next
		c.prevLocationsURL = resource.Previous
		resource.printNames()
	}

	return err
}

func (c *Config) PrintNextListLocationAreas() error {
	return c.printNames(c.nextLocationsURL)
}

func (c *Config) PrintPreviousListLocationAreas() error {

	if c.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	return c.printNames(c.prevLocationsURL)
}
