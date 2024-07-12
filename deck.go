package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// inicjalizacja typu
type deck []string

// Funkcja tworząca stos kart
func newDeck() deck {
	// inicjalizacja type deck slice - przypisanie do zmiennej
	cards := deck{}

	//slice - grow and shrink
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			//dodaje do type
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Funkcja zwracająca 2 slicy
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Funkcja konwertująca - przyjmujące reciver i zwracająca receiver jako string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Funkcja zapisująca karty do pliku
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Funkcja odczytująca z pliku oraz error handling
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		//Print error and quit program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	slices := strings.Split(string(bs), ",")
	return deck(slices)
}

// Funkcja shuffle karty - wykorzystuje random generator
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// funkcja printujaca karty
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
