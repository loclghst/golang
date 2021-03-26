package main

func main() {
	// var card string = "Ace of spades" // long way of declaring vars

	// we can also write

	// var card = "Ace of spades" // a bit short format, type is inferred by the compiler

	// or

	// card := "Ace of spades" // declare and initialise, type is inferred by compiler

	cards := deck{"Ace of Spades", newCard()} // slice of cards

	// append doesn't overwrite the original cards,
	// it returns a new copy after adding the data
	cards = append(cards, "Queen of Diamonds") // add to the cards collection

	// iterate over the collection
	// for i, card := range cards {
	// 	we use := cause for every element new i, cards are declared and initialised
	// 	fmt.Println(i, card)
	// }

	cards.print()

}

func newCard() string {
	return "King of Hearts"
}
