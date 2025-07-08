package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Waterbootdev/pokedex/internal/pokeapi"
)

func PrintPrompt() {
	fmt.Print("Pokedex > ")
}

func ScanCleanInput(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return CleanInput(scanner.Text())
}

func PrintCommand(cleanInput []string, config *pokeapi.Config) {

	if !(len(cleanInput) > 0 && HandleCommand(cleanInput[0], config)) {
		fmt.Println("Unknown command")
	}
}

func main() {

	config := pokeapi.InitConfig()

	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	InitRegistry()

	for {

		PrintPrompt()

		PrintCommand(ScanCleanInput(scanner), &config)
	}
}
