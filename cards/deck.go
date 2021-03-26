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

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "King", "Queen", "Jack"}

	// if we are not using the index during the iteration, we replace
	// it with a _ to avoid the `declared but not used` error

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
