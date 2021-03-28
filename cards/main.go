package main

import "fmt"

func main() {
	// var card string = "Ace of spades" // long way of declaring vars

	// we can also write

	// var card = "Ace of spades" // a bit short format, type is inferred by the compiler

	// or

	// card := "Ace of spades" // declare and initialise, type is inferred by compiler

	// cards := []string{"Ace of Spades", newCard()} // slice of cards

	// cards := deck{"Ace of Spades", newCard()} // using custom type

	// append doesn't overwrite the original cards,
	// it returns a new copy after adding the data
	// cards = append(cards, "Queen of Diamonds") // add to the cards collection

	// iterate over the collection

	// for i, card := range cards {
	// 	we use := cause for every element in cards,
	//  i, cards are re-declared and re-initialised

	// 	fmt.Println(i, card)
	// }

	cards := newDeck()

	//getting multiple return values

	hand, remainingCards := deal(cards, 4)

	// we can call print() on hand and remainingCards cause both are of type deck
	hand.print()
	remainingCards.print()

	// greeting := "Hi there"

	// type casting ----> type(valueToConvert)
	// byteSlice := []byte(greeting)

	// fmt.Println(byteSlice)

	// fmt.Println(cards.toString())

	cards.saveToFile("my_cards")

	fmt.Println("***************")
	// newDeckFromFile("my_cards").print()

	cards.shuffle()

	fmt.Println(cards)

}

func newCard() string {
	return "King of Hearts"
}
