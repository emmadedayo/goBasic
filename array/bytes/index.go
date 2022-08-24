package main

import (
	"io/ioutil"
)

func main() {

	cards := cardDec()
	err := cards.saveFunction("mycard")
	if err != nil {
		return
	}
	//fmt.Println(cards.toString())

	greetings := "Hello Emmanuel Adenagbe"

	ioutil.WriteFile("Test", []byte(greetings), 777)
}
