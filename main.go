package main

import "fmt"

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func commandHelp() error {
	return nil
}

func commandExit() error {
	return nil
}

func main() {

	// create a map for available commands, every command is an instance of that struct
	Commands := map[string]CliCommand{
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

	fmt.Printf("%+v", Commands)

	// stdin
	// inifite loop, break if the stdin has one of the commands name
	// execute the callback of that command

	println("Hello")
}
