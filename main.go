package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, value := range GetCommands() {
		fmt.Printf("%s: %s\n", value.Name, value.Description)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func GetAllKeys(m map[string]CliCommand) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading input: %+v", err)
		}

		Commands := GetCommands()
		command := strings.ToLower(scanner.Text())
		words := strings.Fields(command)

		if len(words) == 0 {
			continue
		}

		cmd, ok := Commands[words[0]]

		if ok {
			cmd.Callback()
		} else {
			fmt.Printf("Command unrecognized, try one of those %v\n", GetAllKeys(Commands))
		}

	}

}
