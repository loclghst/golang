package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 160 {
		t.Errorf("Expected deck of length 16 but instead got %v", len(d))
	}
}
