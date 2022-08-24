package main

import "fmt"

var variableString string

func main() {

	variableString = "i'm Emmanuel Adenagbe"
	fmt.Println(variableString)

	var variable int
	variable = 10
	fmt.Println(variable)

	variableString = "i'm Emmanuel Adenagbe"
	fmt.Println(variableString)

	variableString = newString()
	fmt.Println(variableString)
}

func newString() string {
	return "new string from a function"
}
