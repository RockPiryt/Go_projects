package main

import (
	"fmt"
	"math"
)

// Obliczanie pierwiastka kwadratowego
func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("Nie można obliczyć pierwiastka z liczy ujemnej!")
	}
	return math.Sqrt(number), nil
}

func main() {
	root, err := squareRoot(4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.3f", root)
	}
}
