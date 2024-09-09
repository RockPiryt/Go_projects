package main

import "fmt"

type book string // classa ksiązka wraz z typem

//w nawiasie jaki argument (b) jakiej clasy(book)
func (b book) printTitle() {
	fmt.Println(b)
}

func main() {
	var b book = "Harry Potter"
	b.printTitle()
}
