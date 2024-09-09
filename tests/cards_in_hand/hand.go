package main

import "fmt"

//------------------------------------CLASA I METODA

// tworze klase reka która ma typ listy
type hand []string

//tworzę metodę dla klasy
func (h hand) printAllCards() {
	//wyprintowanie po kolei
	for i, card := range h {
		fmt.Println(i, card)
	}
}

//----------------------------------------FUNKCJE

// Funkcja tworząca stos kart w ręce
func createHand() hand {

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
	return myHand
}

//funkcja zwracajaca 2 wartości - to nie jest metoda
//Funkca dzieli listę na 2 częsci

func podzielDwaStosy(h hand) (hand, hand) {
	return h[:3], h[3:]
}
