package main

import "fmt"

// tworze klase reka która ma typ listy
type hand []string

func main() {
	//tworzę instancję ręki
	myHand := hand{}

	//rozdawane karty
	firstCard := "2 of Spades"
	var secondCard string = "3 of Spades"

	//dodanie do clasy, która ma typ listy
	myHand = append(myHand, firstCard, secondCard)

	//wyprintowanie całej listy
	fmt.Println(myHand)

	//wyprintowanie po kolei
	for i, card := range myHand {
		fmt.Println(i, card)
	}

}
