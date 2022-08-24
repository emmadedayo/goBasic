package main

import (
	"os"
	"testing"
)

func TestCartType(t *testing.T) {

	c := cardDec()
	if len(c) != 16 {
		t.Error("Expected 16 cards, but got", len(c))
	}

	if c[0] != "Ace of Spades" {
		t.Error("Expected Ace of Spades, but got", c[0])
	}
}

func TestSaveToFile(t *testing.T) {
	os.Remove("_testDeck")
	c := cardDec()
	err := c.saveFunction("_testDeck")
	if err != nil {
		return
	}

	loadFile := readFile("_testDeck")
	if len(loadFile) != 16 {
		t.Error("Expected 16 cards, but got", len(loadFile))
	}

}
