package main

import (
	"fmt"
	"os"

	"github.com/Waterbootdev/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config) error
}

var registry map[string]cliCommand = make(map[string]cliCommand)

func commandExit(_ *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *pokeapi.Config) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range registry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandMapPrevious(config *pokeapi.Config) error {
	return config.PrintPreviousListLocationAreas()
}

func commandMapNext(config *pokeapi.Config) error {
	return config.PrintNextListLocationAreas()
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

func HandleCommand(command string, config *pokeapi.Config) bool {
	cmd, ok := registry[command]
	if ok {
		err := cmd.callback(config)
		if err != nil {
			fmt.Println(err)
		}
	}
	return ok
}
