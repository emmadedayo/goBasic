package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	_ "strings"
	"time"
)

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

func (t cardType) toString() string {
	return strings.Join([]string(t), ",")
}

func (t cardType) saveFunction(fileName string) error {

	return ioutil.WriteFile(fileName, []byte(t.toString()), 777)

}

func readFile(filename string) cardType {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}
	ss := strings.Split(string(bs), ",")

	return cardType(ss)
}

func (t cardType) random() {

	ss := rand.NewSource(time.Now().UnixNano())
	r := rand.New(ss)

	for i := range t {
		newPosition := r.Intn(len(t) - 1)
		t[i], t[newPosition] = t[newPosition], t[i]
	}
}
