package main

import "fmt"

func main() {

	// Stworzenie maps
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "459gh22",
	}

	//Inicjalizacja bez wartości - sposób 1
	var colorsA map[string]string
	//Inicjalizacja bez wartości - sposób 2
	colorsB := make(map[string]string)

	// Dodanie albo update
	colors["white"] = "#ffffff"

	// Delete po key
	delete(colors, "red")

	//Użycie funkcji do printowania
	printMap(colors)

	fmt.Println(colorsA)
	fmt.Println(colorsB)

}

//Iterowanie po maps
func printMap(col map[string]string) {
	for color, hex := range col {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
