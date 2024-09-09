package main

import "fmt"

func main() {
	//tworzę instancję ręki
	myHand := hand{}

	//rozdawane karty
	firstCard := "2 of Spades"
	var secondCard string = "3 of Spades"
	thirdCard := "4 of Diamonds"

	//dodanie do instancji clasy (typ listy)
	myHand = append(myHand, firstCard, secondCard, thirdCard)

	//wyprintowanie całej listy
	fmt.Println(myHand)

	//zastosowanie metody klasy hand
	myHand.printAllCards()

}
