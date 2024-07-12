package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	mydeck := newDeck()

	// Check number of card if is equal to 16
	if len(mydeck) != 16 {
		t.Errorf("Expected deck length of 16, but go %v", len(mydeck))
	}

	// // Fail
	// if len(mydeck) != 200 {
	// 	t.Errorf("Expected deck length of 200, but go %v", len(mydeck)) //got 16
	// }

	// Check first card name
	if mydeck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but go %v", mydeck[0])
	}

	// Check last card name
	if mydeck[len(mydeck)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Ace of Spades, but go %v", mydeck[len(mydeck)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	//Remove previous files
	os.Remove("_decktesting")

	// Create new file with deck
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

}
