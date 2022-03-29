package main

import "fmt"

func main() {

	// Ways to declare maps.
	// var colors map[string]string
	// colors := make(map[string]string)

	// := declaring and initializing.

	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"black": "#000000",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
