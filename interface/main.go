package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	printGreetingInEnglish(englishBot{})
	printGreetingInSpanish(spanishBot{})
}

func printGreetingInEnglish(b bot) {
	fmt.Println(b.getGreeting())
}

func printGreetingInSpanish(sp spanishBot) {
	fmt.Println(sp.getGreeting())
}

func (eg englishBot) getGreeting() string {
	return "Hi there!"
}

func (sp spanishBot) getGreeting() string {
	return "Hola!"
}
