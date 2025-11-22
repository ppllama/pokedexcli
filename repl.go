package main

import(
	"strings"
	"bufio"
	"os"
	"fmt"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}


func StartRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		command := words[0]

		com, ok := getCommands()[command]

		if !ok {
			fmt.Printf("Unknown command\n")
			continue
		}

		if err := com.callback(conf); err != nil {
			fmt.Printf("error in callback: %s\n", err)
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	split := strings.Fields(lower)
	return split
}


func getCommands() map[string]cliCommand {
	registry := map[string]cliCommand{
		"exit": {
			name:		"exit",
			description:"Exit the Pokedex",
			callback:	commandExit,
		},
		"help": {
			name:		"help",
			description:"Displays a help message",
			callback: 	commandHelp,
		},
		"map": {
			name:		"map",
			description:"Lists next 20 location areas",
			callback:	commandMap,
		},
		"mapb": {
			name:		"mapb",
			description:"Lists previous 20 location areas",
			callback: 	commandMapb,
		},
	}
	return registry
}