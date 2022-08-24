package main

import "fmt"

func main() {

	var numberA, numberB int

	fmt.Println("Welcome to the program to add two numbers")
	fmt.Println("Enter the first number")
	fmt.Scanln(&numberA)
	fmt.Println("Good,Enter the second number")
	fmt.Scanln(&numberB)
	sum := numberA + numberB
	fmt.Println("Sum of two numbers is", sum)
}
