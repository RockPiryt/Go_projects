package main

import "fmt"

//Określenie interface
type bot interface {
	getGreeting() string

	// Inne przykłady przypisania funkcji
	// getGreeting(string, int) (string,error)
	// getBotVersion() float64
	// respondToUser() string

}

//Określenie structów
type englishBot struct{}
type spanishBot struct{}

func main() {
	//Inicjalizacja structów
	eb := englishBot{}
	sb := spanishBot{}

	//Wykorzystanie prostej zbliżonej funkcji
	printGreeting(eb)
	printGreeting(sb)

}

// --------------------Jedna wspólna funkcja
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//------------------------------Zbliżone funkcje
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting()) // printowanie wyniku skomplikowanej funkcji
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

//-------------------------------Skomplikowane funkcje - różniące się (bez argumentu)
func (englishBot) getGreeting() string {
	return "Hi There"
}

// func (eb englishBot) getGreeting() string{
// 	return "Hi There"
// }

func (spanishBot) getGreeting() string {
	return "Hola!"
}

// func (sb spanishBot) getGreeting() string{
// 	return "Hola!"
// }
