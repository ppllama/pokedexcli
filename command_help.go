package main

import(
	"fmt"
)

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n")
	fmt.Print("\n")
	for registered := range(getCommands()) {
		fmt.Printf("%s: %s\n", getCommands()[registered].name, getCommands()[registered].description)
	}
	return nil
}