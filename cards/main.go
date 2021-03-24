package main

import "fmt"

func main() {
	// var card string = "Ace of cards"
	cards := []string{"Ace of Spades", newCard()} // slice of cards

	cards = append(cards, "Queen of Diamonds") // add to the cards collection

	for i, card := range cards { //iterate over the collection
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "King of Hearts"
}