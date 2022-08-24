package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}
type ContactInfo struct {
	email   string
	phone   string
	zipCode int
}

func main() {

	p := Person{"John", "Doe", ContactInfo{
		email:   "joedoe@gmail.com",
		phone:   "1234567890",
		zipCode: 12345,
	}}
	fmt.Println(p)

	//animalName := AnimalName{"John", "Doe"}
	//fmt.Println(animalName)
	//
	//name := name()
	//fmt.Println(name)
	//
	//fmt.Println(name.print())

	//newNamePointer := &p
	p.updatePersonName("Emmanuel")
	p.print()
	fmt.Println()

}

func (an *Person) updatePersonName(name string) {
	(*an).firstName = name
}

func (an Person) print() {
	fmt.Printf("%+v", an)
}
