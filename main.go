package main

import (
	"bufio"
	"fmt"
	"os"
)

func PrintPrompt() {
	fmt.Print("Pokedex > ")
}

func ScanCleanInput(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return CleanInput(scanner.Text())
}

func PrintCommand(cleanInput []string, config *config) {

	if !(len(cleanInput) > 0 && HandleCommand(cleanInput[0], config)) {
		fmt.Println("Unknown command")
	}
}

func main() {

	config, err := InitConfig(20)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	InitRegistry()

	for {

		PrintPrompt()

		PrintCommand(ScanCleanInput(scanner), &config)
	}
}
