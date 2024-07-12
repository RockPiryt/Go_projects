package main

func main() {

	//cards := newDeck()

	// cards.print()
	//-----------------podział na 2 slicy
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	//-------------------zmiana deck do []byte
	// fmt.Println((cards.toString()))

	//------------------zapisanie do pliku
	//cards.saveToFile("bla")

	//-----------------odczyt tali z pliku
	// cards := newDeckFromFile("my_file_cards")
	// cards.print()

	//-------------shuffle
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
