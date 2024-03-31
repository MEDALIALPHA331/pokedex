package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func runPokeRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading input: %+v", err)
		}

		Commands := GetCommands()

		words := cleanCommandText(scanner.Text())

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
		"map": {
			Name:        "map",
			Description: "",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "",
			Callback:    commandMapb,
		},
	}
}
