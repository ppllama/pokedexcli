package main

import(
	"strings"
	// "fmt"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	// fmt.Print(lower)
	split := strings.Split(lower, " ")
	var output []string
	for _, word := range(split) {
		if word != "" {
			output = append(output, word)
		}
	}
	return output
}