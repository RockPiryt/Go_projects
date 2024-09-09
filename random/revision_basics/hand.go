package main

import "fmt"

// tworze klase reka która ma typ listy
type hand []string

//tworzę metodaę dla klasy
func (h hand) printAllCards() {
	//wyprintowanie po kolei
	for i, card := range h {
		fmt.Println(i, card)
	}
}

//funkcja zwracajaca 2 wartości
//Funkca dzieli listę na 2 częsci

func podzielDwaStosy(h hand) (hand, hand) {
	return h[:3], h[3:]
}
