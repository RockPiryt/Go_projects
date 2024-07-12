package main

import "fmt"

func main() {
	c := color("Red") //tworze isntancje typu color

	fmt.Println(c.describe("is an awesome color"))
}

//Typ
type color string

//Funkcja do typu
func (c color) describe(description string) string {
	return string(c) + " " + description
}
