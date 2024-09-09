package main

import (
	"fmt"
)

func main() {
	//Sposoby przypisania
	var oneCard string = "Ace of Spades"
	secondCard := "Five of diamonds"
	fmt.Println(oneCard)
	fmt.Println(secondCard)

	//Tworzę listę
	leftHand := []string{"Six of Spades", "Seven of Spades"}
	fmt.Println(leftHand)

	//dodaje do listy
	fmt.Println("Dodaje 1 karte")
	leftHand = append(leftHand, "8 of spades")
	fmt.Println(leftHand)

	//Dodaje do listy karte za pomocą funckji
	fmt.Println("Dodaje 1 karte za pomoca funkcji")
	leftHand = append(leftHand, "9 of spades", newCard())
	fmt.Println(leftHand)

	//iteruje pod slice, czyli po liscie z numeracja
	fmt.Println("iteruje pod slice, czyli po liscie z numeracja")
	for i, card := range leftHand {
		fmt.Println(i, card)
	}

	// //iteruje pod slice, czyli po liscie z numeracja
	// fmt.Println("iteruje pod slice, czyli po liscie bez numeracji")
	// for _, card := range leftHand {
	// 	fmt.Println( card)
	// }

}

// Funkcja tworzaca karte
func newCard() string {
	return "10 of spades"
}
