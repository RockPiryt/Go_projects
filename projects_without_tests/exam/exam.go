package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Podaj wynik: ")
	reader := bufio.NewReader(os.Stdin)   //ma pobierać dane wejsciowe z klawiatury
	input, err := reader.ReadString('\n') //wczytuje jako string do momentu naciśniecia enter
	if err != nil {                       // gdyby był error w input to wyswietli error i zakoczy program
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)            //usuwanie białych znaków (entera)
	grade, err := strconv.ParseFloat(input, 64) //przekształcenie stringa na float
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "zaliczenie"
	} else {
		status = "brak zaliczenia"
	}
	fmt.Println("Wynik", grade, "oznacza", status, ".")

}
