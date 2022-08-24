package main

import "fmt"

func main() {

	numbersValue := []int{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
	}

	for _, numbersValue := range numbersValue {
		if numbersValue%2 == 0 {
			fmt.Println(numbersValue, "is even")
		} else {
			fmt.Println(numbersValue, "is odd")
		}
	}
}
