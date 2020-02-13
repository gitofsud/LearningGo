package main

import "fmt"

func main() {
	// Introduction
	colors := map[string]string {
		"red": "Sun",
		"green": "Trees",
	}

	fmt.Println(colors)

	// Manipulating Map
	cars := make(map[string]string)

	cars["white"] = "Audi"
	cars["black"] = "BMW"

	delete(cars, "black")

	fmt.Println(cars)

	// Iterating Map
	words := map[string]string {
		"x": "Xavier",
		"y": "Yesterday",
		"z": "Zebra",
	}

	printMap(words)
}

func printMap(c map[string]string) {
	for i, x := range c {
		fmt.Println(i, ": ", x)
	}
}