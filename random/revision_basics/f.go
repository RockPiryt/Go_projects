package main

import "fmt"

func main() {
	//tworzę instancję ręki
	myHand := hand{}

	//listy
	literyLista := []string{"A", "B", "C"}
	liczbyLista := []string{"1", "2", "3"}

	//stworznie 1 listy z dwóch
	for _, liczba := range liczbyLista {
		for _, litera := range literyLista {
			myHand = append(myHand, liczba+litera)
		}
	}
	// printowanie całej listy [1A 1B 1C 2A 2B 2C 3A 3B 3C]
	fmt.Println(myHand)

	//wykorzystanie metody clasy hand do printowania po kolei
	myHand.printAllCards()

}
