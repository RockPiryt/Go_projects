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

	// zwracanie wielu wartosci przez funkcje i przypisanie ich do zmiennych
	// podziele całą liste [1A 1B 1C 2A 2B 2C 3A 3B 3C] na 2 stosy
	stosFirst, stosSecond := podzielDwaStosy(myHand)

	fmt.Println(stosFirst)
	fmt.Println(stosSecond)

}
