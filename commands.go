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

func subMapCommand(config *config, locationsResp *pokeapi.Resource) {

	config.nextLocationsURL = locationsResp.Next
	config.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
}

func commandMapb(config *config) error {
	locationResp, err := config.PreviousListLocationAreas()
	if err != nil {
		return err
	}
	subMapCommand(config, locationResp)
	return nil
}

func commandMap(config *config) error {
	locationsResp, err := config.NextListLocationAreas()
	if err != nil {
		return err
	}
	subMapCommand(config, locationsResp)
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
		callback:    commandMap,
	}
	registry["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous 20 location areas",
		callback:    commandMapb,
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
