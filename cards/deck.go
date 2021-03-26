package main

import "fmt"

// custom type declaration
type deck []string

// function with a receiver
// d is a variable which points to the working copy of the deck instance
// make the function available to any variable of type deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
