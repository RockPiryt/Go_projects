package main

import "fmt"

func main() {

	// -------------------------------------------------Sama tworze listę i wykorzytsuje metodę printAllCards clasy
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

	//-------------------------------------------------- funkcja która wykorzystuje classe hand
	myRightHand := createHand()
	fmt.Println("Drugi zestaw kart")
	myRightHand.printAllCards()

	fmt.Println("Drugi zestaw kart ma ilość: ", len(myRightHand))

}
