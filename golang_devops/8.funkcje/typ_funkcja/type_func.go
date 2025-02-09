package main

import "fmt"

//wstępna deklaracja
type MyFunc = func(x,y int) int 

func main() {
	var f MyFunc = func(a,b int) int{
		return a + b
	}

	x := f(2,3)
	fmt.Printf("Funkcja powinna zwrocic 5, a zwraca:%d", x)
}