package main

type Color string

func (c Color) describe(description string) string {
	return string(c) + " " + description
}
