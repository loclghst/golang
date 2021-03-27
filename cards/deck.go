package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// custom type declaration
type deck []string

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

// function with a receiver
// d is a variable which points to the working copy of the deck instance
// make the function available to any variable of type deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// returning multiple values
// func deal returns two values both of type deck
func deal(d deck, handSize int) (deck, deck) {
	// multiple return values are seperated by a comma
	// to divide a slice into two parts we can use
	// slice[startIndexIncluding : uptoNotIncluding]
	// we can also omit both the indexes and compiler will infer correctly
	return d[:handSize], d[handSize:]
}

// convert toString, typecasting

func (d deck) toString() string {
	//strings.Join() is coming from strings package
	return strings.Join([]string(d), ",")
}

// save to file using ioutil package
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		//error handling
		fmt.Println("Error: ", err)

		//non zero exit code means the is an error
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))

}
