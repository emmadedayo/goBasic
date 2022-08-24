package main

import "fmt"

func main() {

	//colors := map[string]string{
	//	"red":   "#ff0000",
	//	"green": "#00ff00",
	//}

	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"
	colors["blue"] = "#0000ff"
	colors["yellow"] = "#ffff00"

	//delete(colors, "red")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
