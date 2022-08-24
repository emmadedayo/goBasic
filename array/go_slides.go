package main

import "fmt"

func main() {

	//Creating a string of array
	card := []string{
		"Hello world",
		"Hello world",
	}
	card = append(card, "Hello Emmanuel")

	fmt.Println(card[3:])

	for _, card := range card {
		fmt.Println(card)
	}
}
