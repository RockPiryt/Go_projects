package mystring

// Reverse odwraca string znak po znaku
func Reverse(s string) string {
    runes := []rune(s)  // Konwertujemy string na tablicę run (obsługa znaków Unicode)
    length := len(runes) // Długość tablicy

    for i := 0; i < length/2; i++ {
        j := length - 1 - i // Indeks przeciwległego znaku
        runes[i], runes[j] = runes[j], runes[i] // Zamieniamy miejscami znaki
    }

    return string(runes) // Konwertujemy tablicę run na string i zwracamy wynik
}

// String w Go jest sekwencją bajtów, a rune to jeden znak Unicode.
// Dzięki []rune(s) obsługujemy też polskie litery i inne znaki spoza ASCII.