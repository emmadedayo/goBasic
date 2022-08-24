package main

type AnimalName struct {

	// The name of the animal
	MaleName   string
	FemaleName string
}

func name() AnimalName {
	animal := AnimalName{}

	animal.MaleName = "Bull"
	animal.FemaleName = "Goat"
	return animal
}

func (an AnimalName) print() string {
	return an.MaleName + an.FemaleName
}
