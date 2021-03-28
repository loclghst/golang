package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck of length 16 but instead got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but instead got %v", d[0])
	}

	if d[len(d)-1] != "Jack of Hearts" {
		t.Errorf("Expected last card to be Jack of Hearts but instead got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// deletes a file or folder of given name
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck of length 16 but instead got %v", len(loadedDeck))
	}
	// clean up
	os.Remove("_decktesting")

}
