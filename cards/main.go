package main

func main() {
	// var card string = "Ace of cards"
	cards := deck{"Ace of Spades", newCard()} // slice of cards

	cards = append(cards, "Queen of Diamonds") // add to the cards collection

	// for i, card := range cards { //iterate over the collection
	// 	fmt.Println(i, card)
	// }

	cards.print()

}

func newCard() string {
	return "King of Hearts"
}
