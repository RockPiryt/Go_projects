package main

import (
	"fmt"
	"strconv"
)

func main(){
	v := "1"
	if a, err := strconv.Atoi(v); err == nil {
		fmt.Printf(" typ %T, wartość %d\n", a,a)
	}

	//fmt.Println(a) a nie jest dostępne poza ifem

	//ponowne wykorzystanie zminenej dlatego =
	//chce wyłapać błąd - zalecane
	v = "3232656565545432132154548797542463"

	if _, err := strconv.Atoi(v); err !=nil {
		fmt.Printf("Funkcja zwróciła bład %v\n", err)
	}

	
}