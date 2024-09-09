package main

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

	myHand.printAllCards()

}
