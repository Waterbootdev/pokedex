package main

import (
	"fmt"
	"os"

	"github.com/Waterbootdev/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

var registry map[string]cliCommand = make(map[string]cliCommand)

func commandExit(_ *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *config) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range registry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandMap(config *config, getListLocationAreas func() (pokeapi.Resource, error)) error {
	resource, err := getListLocationAreas()

	if err != nil {
		return err
	}

	config.nextLocationsURL = resource.Next
	config.prevLocationsURL = resource.Previous

	for _, loc := range resource.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapPrevious(config *config) error {
	return commandMap(config, config.PreviousListLocationAreas)
}

func commandMapNext(config *config) error {
	return commandMap(config, config.NextListLocationAreas)
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
}

func HandleCommand(command string, config *config) bool {
	cmd, ok := registry[command]
	if ok {
		err := cmd.callback(config)
		if err != nil {
			fmt.Println(err)
		}
	}
	return ok
}
