package pokeapi

import (
	"errors"
	"time"
)

type Config struct {
	client                   Client
	previousLocationAreasURL *string
	nextLocationAreasURL     *string
}

func InitConfig() Config {
	pokeClient := NewClient(5 * time.Second)
	return Config{client: pokeClient}
}

func LocationAreasURL(url *string) string {
	if url == nil {
		return locationAreasURL
	} else {
		return *url
	}
}

func (c *Config) printNames(locationAreasURL *string) error {

	locationAreasResource, err := c.client.ListLocationAreas(LocationAreasURL(locationAreasURL))

	if err != nil {
		return err
	}

	c.previousLocationAreasURL = locationAreasResource.Previous
	c.nextLocationAreasURL = locationAreasResource.Next

	locationAreasResource.printNames()

	return err
}

func (c *Config) PrintNextListLocationAreas() error {

	return c.printNames(c.nextLocationAreasURL)
}

func (c *Config) PrintPreviousListLocationAreas() error {

	if c.previousLocationAreasURL == nil {
		return errors.New("you're on the first page")
	}

	return c.printNames(c.previousLocationAreasURL)
}
