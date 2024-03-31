package main

import (
	"strings"
)

func GetAllKeys(m map[string]CliCommand) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func cleanCommandText(command string) []string {
	command = strings.ToLower(command)
	words := strings.Fields(command)
	return words
}
