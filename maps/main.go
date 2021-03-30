package main

import "fmt"

func main() {
	// declaring a map
	// [string] means keys are of type string, next string means values are of type string
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ff1234",
	}

	fmt.Println(colors)
}
