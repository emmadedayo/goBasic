package main

import "fmt"

func main() {

	//cards := cardType{
	//	"Ace",
	//	"Card",
	//}

	cards := cardDec()
	cards.print()

	slideCard := []int{
		1,
		2,
		3,
		4,
		5,
		6,
	}

	fmt.Println(slideCard[0:3])

	fmt.Println('\n')

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	title, number := getBookInfo()
	fmt.Println(title)
	fmt.Println(number)

	colorA, colorB, colorC := color()
	fmt.Println(colorA, colorB, colorC)
}

func getBookInfo() (string, int) {
	return "Hello", 100
}

func color() (string, string, string) {
	return "red", "blue", "yellow"
}
