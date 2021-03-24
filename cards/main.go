package main

import "fmt"

func main() {
	// var card string = "Ace of cards"
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Ace of Spades"
}
