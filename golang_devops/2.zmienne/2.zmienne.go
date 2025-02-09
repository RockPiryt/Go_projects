package main

import "fmt"

func main() {
	var A int64 = 5
	var B int32
	C := 1

	A = int64(C)

	var D string = "Ala ma kotka"
	var E string
	F := "Kot ma ale"
	fmt.Println("Zmienna A jest rowna ", A)
	fmt.Println("Zmienna B jest rowna ", B)
	fmt.Println("Zmienna C jest rowna ", C)

	fmt.Printf("Zmienna A jest typu %T\n ", A)
	fmt.Printf("Zmienna B jest typu %T\n ", B)
	fmt.Printf("Zmienna C jest typu %T\n ", C)
	fmt.Printf("Zmienna D jest typu %T\n ", D)
	fmt.Printf("Zmienna E jest typu %T\n ", E)
	fmt.Printf("Zmienna F jest typu %T\n ", F)

}
