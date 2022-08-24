package main

import (
	"fmt"
	"strconv"
)

func main() {

	var parameter = lengthOfRectangle(900, 5)
	fmt.Printf("length of rectangle is %s", parameter)

	fmt.Printf("\n")

	var area = areaOfATriangle(10, 5)
	fmt.Printf("area of a triangle is %d", area)

	fmt.Printf("\n")
	var areaOfTrapezium = areaOfTrapezium(10, 5, 5)
	fmt.Printf("area of a trapezium is %d", areaOfTrapezium)
}

func lengthOfRectangle(parameter, width int) string {

	if parameter < width {
		return "parameter is less than width"
	} else {
		// using formula p/2 - w
		var length = parameter/2 - width
		//strconv is to convert from int to string
		return strconv.Itoa(length)
	}
}

func areaOfATriangle(base, height int) int {

	var area = (base * height) / 2
	return area
}

func areaOfTrapezium(aBase, bBase, Height int) int {
	var area = ((aBase + bBase) * Height) / 2
	return area
}
