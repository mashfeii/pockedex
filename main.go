package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mashfeii/pokedexcli/internal/api"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Magenta = "\033[35m"

const cliName = "pokedex"

func printPrompt() {
	fmt.Print(cliName, " > ")
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func handleUnknown(command string) {
	fmt.Printf("Unknown command: %s | Try 'help' to see available commands\n", command)
}

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func displayHelp() {
	fmt.Println("Welcome to PockedexCli! These are the list of available commands:")
	commands := possibleCommands()
	for _, value := range commands {
		fmt.Printf("%v%s%v - %s\n", Magenta, value.name, Reset, value.description)
	}
	fmt.Println()
}

var config = api.Config{
	Next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	Previous: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
}

func possibleCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays help message",
			callback:    displayHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pockedex",
		},
		"map": {
			name:        "map",
			description: "Show the next 20 locations",
			callback:    func() { api.GetLocations(&config, true) },
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 locations",
			callback:    func() { api.GetLocations(&config, false) },
		},
	}
}

func main() {
	commands := possibleCommands()
	scanner := bufio.NewScanner(os.Stdin)
	printPrompt()
	for scanner.Scan() {
		text := cleanInput(scanner.Text())

		if strings.EqualFold(text, "exit") {
			fmt.Println("Gracefully shutting down...")
			return
		} else if _, finded := commands[text]; finded {
			commands[text].callback()
		} else {
			handleUnknown(text)
		}

		printPrompt()
	}
}
