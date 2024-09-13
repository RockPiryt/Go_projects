package main

import (
	"fmt"
	"strings"
)

func main() {
	phrases := []string("moja siostra", "kot")
	fmt.Println("Zdjęcie na którym są", JoinWithCommas(phrases))

	phrases = []string("mama", "tata", "Leon")
	fmt.Println("Zdjęcie na którym są", JoinWithCommas(phrases))
}

func JoinWithCommas(phrases []string) string {
	result := strings.Join(phrases[:len(phrases)-1], ",")
	result += " i "
	result += phrases[len(phrases)-1]
	return result
}
