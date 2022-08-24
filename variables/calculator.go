package main

import "fmt"

func main() {

	var d int
	var e float32

	a := 10
	b := 10
	c := a + b
	fmt.Println(c)

	d = calculator(a, b)
	fmt.Println(d)

	e = float32(areaOfACircle(10.0))
	fmt.Println(e)

}

func calculator(a, b int) int {
	return a + b
}

func areaOfACircle(radius float64) float64 {
	return 3.14 * radius * radius
}
