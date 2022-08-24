package main

import "fmt"

type cardType []string

func cardDec() cardType {
	cards := cardType{}

	cardsName := []string{
		"Spades",
		"Diamond",
		"Hearts",
		"Club",
	}

	cardsValue := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
	}

	for _, cardsName := range cardsName {
		for _, cardsValue := range cardsValue {
			cards = append(cards, cardsValue+" of "+cardsName)
		}
	}
	return cards
}

func (t cardType) print() {
	for _, cardType := range t {
		fmt.Println(cardType)
	}
}

func deal(t cardType, handSize int) (cardType, cardType) {
	return t[:handSize], t[:handSize]
}
