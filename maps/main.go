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

	// we can also declare maps as follows
	var newMap map[string]string // this is initialised with Zero values
	fmt.Println(newMap)

	// another way of declaring map
	anotherMap := make(map[string]string) // this is also initialsed with zero values

	fmt.Println(anotherMap)

}
