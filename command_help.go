package main

import "fmt"

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, value := range GetCommands() {
		fmt.Printf("%s: %s\n", value.Name, value.Description)
	}
	return nil
}
