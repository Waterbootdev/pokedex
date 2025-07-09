package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Waterbootdev/pokedex/internal/pokeapi"
	"github.com/Waterbootdev/pokedex/internal/pokecache"
)

func PrintPrompt() {
	fmt.Print("Pokedex > ")
}

func ScanCleanInput(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return CleanInput(scanner.Text())
}

func PrintCommand(cleanInput []string, config *pokeapi.Config) {

	ok := false

	switch len(cleanInput) {

	case 1:
		ok = HandleCommand(cleanInput[0], config, "")

	case 2:
		ok = HandleCommand(cleanInput[0], config, cleanInput[1])

	}

	if !ok {
		fmt.Println("Unknown command")
	}
}

func main() {

	config := pokeapi.InitConfig(pokecache.NewCache(5 * time.Second))

	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	InitRegistry()

	for {

		PrintPrompt()

		PrintCommand(ScanCleanInput(scanner), &config)
	}
}
