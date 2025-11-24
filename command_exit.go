package main

import(
	"os"
	"fmt"
)


func commandExit(_ *config, _ string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}