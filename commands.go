package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Waterbootdev/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config, args ...string) error
}

var registry map[string]cliCommand = make(map[string]cliCommand)

func commandExit(_ *pokeapi.Config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *pokeapi.Config, _ ...string) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range registry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandMapPrevious(config *pokeapi.Config, _ ...string) error {
	return config.PrintPreviousListLocationAreas()
}

func commandMapNext(config *pokeapi.Config, _ ...string) error {
	return config.PrintNextListLocationAreas()
}

func commandExplore(config *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no location area provided")
	}

	locationAreaName := args[0]

	fmt.Println("Exploring", locationAreaName, "...")
	fmt.Println("Found Pokemon:")
	fmt.Println()

	return config.PrintPokemonInLocationArea(locationAreaName)
}

func commandCatch(config *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	fmt.Println("Throwing a Pokeball at", pokemonName+"...")

	caught, err := config.PrintPokemonInLocationAre(pokemonName)

	if err != nil {
		return err
	}

	if caught {
		fmt.Println(pokemonName, "was caught!")
	} else {
		fmt.Println(pokemonName, "escaped!")
	}

	return nil
}

func commandInspect(config *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	return config.InspectPokemon(pokemonName)
}

func commandPokedex(config *pokeapi.Config, _ ...string) error {
	fmt.Println("Your Pokedex:")
	fmt.Println()
	config.PrintCaughtPokemon()
	return nil
}
func InitRegistry() {

	registry["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	registry["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	registry["map"] = cliCommand{
		name:        "map",
		description: "Displays the next 20 location areas",
		callback:    commandMapNext,
	}
	registry["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous 20 location areas",
		callback:    commandMapPrevious,
	}
	registry["explore"] = cliCommand{
		name:        "explore",
		description: "Explore a location area",
		callback:    commandExplore,
	}
	registry["catch"] = cliCommand{
		name:        "catch",
		description: "Attempt to catch a pokemon",
		callback:    commandCatch,
	}
	registry["inspect"] = cliCommand{
		name:        "inspect",
		description: "Inspect a pokemon",
		callback:    commandInspect,
	}
	registry["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "displays all caught Pokemons",
		callback:    commandPokedex,
	}

}

func HandleCommand(command string, config *pokeapi.Config, argument string) bool {
	cmd, ok := registry[command]
	if ok {
		err := cmd.callback(config, argument)
		if err != nil {
			fmt.Println(err)
		}
	}
	return ok
}
