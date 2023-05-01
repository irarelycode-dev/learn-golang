package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 12 {
		t.Errorf("Expected a length of 12 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades , but got %v", d[0])
	}

}

func TestSaveDeckToFileAndReadDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	deck := newDeck()
	deck.saveToFile("_deckTesting")
	loadedDeck := readDeckFromFile("_deckTesting")
	if len(loadedDeck) != 12 {
		t.Errorf("Expected 12 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_deckTesting")
}
