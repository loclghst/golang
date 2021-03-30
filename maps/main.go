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
	var newMap map[string]string // this syntax declares a nil map

	// The default zero value of a map is nil.
	// A nil map is equivalent to an empty map except that elements can’t be added.

	// if we try to add values to newMap, we will get an error

	fmt.Println(newMap)

	// another way of declaring map
	anotherMap := make(map[string]string) // this is also initialsed with zero values

	// adding key value pairs to Map
	anotherMap["ABC"] = "XYZ" // note the [] syntax
	fmt.Println(anotherMap)

	// deleting key value pairs
	delete(anotherMap, "ABC")
	fmt.Println(anotherMap)

}
