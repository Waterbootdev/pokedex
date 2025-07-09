package pokeapi

import (
	"errors"
	"fmt"
	"time"

	"github.com/Waterbootdev/pokedex/internal/pokecache"
)

type Config struct {
	client                   Client
	previousLocationAreasURL *string
	nextLocationAreasURL     *string
	caughtPokemon            map[string]*Pokemon
}

func InitConfig(cache pokecache.Cache) Config {
	pokeClient := NewClient(5*time.Second, cache)
	return Config{
		client:        pokeClient,
		caughtPokemon: make(map[string]*Pokemon),
	}
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

func (c *Config) PrintPokemonInLocationArea(locationAreaName string) error {

	locationAreaURL := locationAreasURL + "/" + locationAreaName

	locationArea, err := c.client.LocationArea(locationAreaURL)

	if err != nil {
		return err
	}

	locationArea.printNames()

	return nil
}

func (c *Config) PrintPokemonInLocationAre(pokemonName string) (bool, error) {

	pokemon, err := c.client.Pokemon(pokemonURL + pokemonName)

	if err != nil {
		return false, err
	}

	caught := pokemon.caught()

	if caught {
		c.caughtPokemon[pokemonName] = pokemon
	}

	return caught, nil
}

func (c *Config) InspectPokemon(pokemonName string) error {

	pokemon, ok := c.caughtPokemon[pokemonName]

	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("\nType: ")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}

	return nil
}

func (c *Config) PrintCaughtPokemon() {
	for _, pokemon := range c.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
}
