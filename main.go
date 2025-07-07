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

func PrintCommand(cleanInput []string) {

	if !(len(cleanInput) > 0 && HandleCommand(cleanInput[0])) {
		fmt.Println("Unknown command")
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	InitRegistry()

	for {

		PrintPrompt()

		PrintCommand(ScanCleanInput(scanner))
	}
}
