package main

import (
	"bufio"
	"fmt"
	"os"
)

func PrintPrompt() {
	fmt.Print("Pokedex > ")
}

func ScanCleanInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return CleanInput(scanner.Text())
}

func PrintCommand(cleanInput []string) {
	command := ""
	if len(cleanInput) > 0 {
		command = cleanInput[0]
	}
	fmt.Printf("Your command was: %v\n", command)
}

func main() {

	for {
		PrintPrompt()
		cleanInput := ScanCleanInput()
		PrintCommand(cleanInput)
	}
}
