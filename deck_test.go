//https://appliedgo.net/testmain/#:~:text=The%20problem%20is%20that%20the,import%20another%20%E2%80%9Cmain%E2%80%9D%20package.
package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	//lenght of deck
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	//first card

	//last card
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {

		t.Errorf("Expected 16, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
